package repositories

import (
	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"

	lop "github.com/samber/lo/parallel"
)

func (r *Repo) GetWalletLogs(filter *domain.FilterRequest, offset int) ([]*domain.WalletLog, int, error) {
	var walletLogs []*models.WalletLogModel
	var totalCount int64

	query := r.db.Model(&models.WalletLogModel{})

	if filter.UserID != 0 {
		query = query.Where("user_id = ?", filter.UserID)
	}
	if filter.OrderBy != "" {
		query = query.Order(filter.OrderBy + " " + filter.Order)
	}

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(filter.Limit).Find(&walletLogs).Error; err != nil {
		return nil, 0, err
	}

	domainWalletLogs := lop.Map(walletLogs, func(walletLog *models.WalletLogModel, i int) *domain.WalletLog {
		return &domain.WalletLog{
			ID:          walletLog.ID,
			Amount:      walletLog.Amount,
			UserID:      walletLog.UserID,
			Status:      walletLog.Status,
			UpdatedByID: walletLog.UpdatedByID,
			CreatedAt:   walletLog.CreatedAt,
			UpdatedAt:   walletLog.UpdatedAt,
		}
	})

	return domainWalletLogs, int(totalCount), nil
}
