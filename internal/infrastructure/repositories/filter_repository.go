package repositories

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// ฟังก์ชันสำหรับดึงคอลัมน์ทั้งหมดจาก Table
func (r *Repo) getAllColumns(tableName string) ([]string, error) {
	var columns []string
	err := r.db.Raw(`
		SELECT column_name 
		FROM information_schema.columns 
		WHERE table_name = ? AND table_schema = current_schema()
	`, tableName).Scan(&columns).Error
	if err != nil {
		return nil, err
	}
	return columns, nil
}

// Build filter query for search, beginDate, and endDate
func (r *Repo) buildFilterQuery(tableName, search, beginDate, endDate string) (string, []interface{}, error) {
	var conditions []string
	var args []interface{}

	// Fetch columns from the primary table
	columns, err := r.getAllColumns(tableName)
	if err != nil {
		return "", nil, errors.New("failed to get columns")
	}

	// Identify column types to decide how to apply the filter
	columnTypes, err := r.getColumnTypes(tableName)
	if err != nil {
		return "", nil, errors.New("failed to get column types")
	}

	// Add search condition
	if search != "" {
		searchPattern := "%" + strings.ToLower(search) + "%"
		var searchConditions []string
		for _, column := range columns {
			columnType := columnTypes[column]
			switch columnType {
			case "text", "varchar", "char":
				searchConditions = append(searchConditions, "LOWER("+tableName+"."+column+") LIKE ?")
				args = append(args, searchPattern)
			case "integer", "bigint", "smallint":
				if _, err := strconv.Atoi(search); err == nil {
					searchConditions = append(searchConditions, tableName+"."+column+" = ?")
					args = append(args, search)
				}
			case "timestamp", "date":
				if parsedDate, err := time.Parse("2006-01-02", search); err == nil {
					searchConditions = append(searchConditions, tableName+"."+column+" = ?")
					args = append(args, parsedDate)
				}
			default:
				// Skip unsupported column types or handle accordingly
			}
		}

		// Add search condition for companyName and login email
		searchConditions = append(searchConditions, "LOWER(company.company_name) LIKE ? OR LOWER(login.email) LIKE ?")
		args = append(args, searchPattern, searchPattern)

		if len(searchConditions) > 0 {
			conditions = append(conditions, "("+strings.Join(searchConditions, " OR ")+")")
		}
	}

	// Add beginDate condition
	if beginDate != "" {
		begin, err := time.Parse("2006-01-02", beginDate)
		if err != nil {
			return "", nil, errors.New("failed to parse beginDate")
		}
		conditions = append(conditions, tableName+".created_at >= ?")
		args = append(args, begin)
	}

	// Add endDate condition
	if endDate != "" {
		end, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			return "", nil, errors.New("failed to parse endDate")
		}
		// Adjust endDate to end of day
		end = end.Add(24*time.Hour - time.Second)
		conditions = append(conditions, tableName+".created_at <= ?")
		args = append(args, end)
	}

	// Add is_delete condition (must be NULL)
	conditions = append(conditions, tableName+".is_delete IS NULL")

	conditionStr := strings.Join(conditions, " AND ")
	return conditionStr, args, nil
}




func (r *Repo) getColumnTypes(tableName string) (map[string]string, error) {
	columnTypes := make(map[string]string)
	rows, err := r.db.Raw(`
        SELECT column_name, data_type 
        FROM information_schema.columns 
        WHERE table_name = ? AND table_schema = current_schema()`, tableName).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var columnName string
		var dataType string
		if err := rows.Scan(&columnName, &dataType); err != nil {
			return nil, err
		}
		columnTypes[columnName] = dataType
	}

	return columnTypes, nil
}

