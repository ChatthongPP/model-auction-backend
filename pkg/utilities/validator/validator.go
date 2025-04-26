package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unsafe"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type fieldError struct {
	err validator.FieldError
}

type ValidationError struct {
	Code    string `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

/*
 *	Validate
 */
func Validate(input interface{}) error {
	validate = validator.New()
	_ = validate.RegisterValidation("date", IsDate)
	_ = validate.RegisterValidation("is-boolean", isBoolean)
	_ = validate.RegisterValidation("isThaiOrEnglish", validateThaiOrEnglish)
	_ = validate.RegisterValidation("isASCII", validateASCII)
	_ = validate.RegisterValidation("isComplexPassword", validateComplexPassword)
	return validate.Struct(input)
}

/*
 *	Format validation rrrors
 */
func FormatValidationErrors(err error) ErrResponse {
	errs := []ValidationError{}
	for _, err := range err.(validator.ValidationErrors) {
		validationErr := ValidationError{
			Code:    strings.ToUpper(err.ActualTag()),
			Field:   err.Field(),
			Message: fieldError{err}.String(),
		}

		errs = append(errs, validationErr)
	}

	serializedErr := ErrResponse{
		Code:    ErrorStatusCode[ERR_VALIDATION_FAILED],
		Message: ErrorMessage[ERR_VALIDATION_FAILED],
		Errors:  errs,
	}

	return *(*ErrResponse)(unsafe.Pointer(&serializedErr))
}

/*
 *	Function to check parameter pattern for valid date
 */
func IsDate(fl validator.FieldLevel) bool {
	regexString := `^\d{4}\-(0?[1-9]|1[012])\-(0?[1-9]|[12][0-9]|3[01])$`
	Regex := regexp.MustCompile(regexString)
	return Regex.MatchString(fl.Field().String())
}

/*
 *	Function to check boolean
 */
func isBoolean(fl validator.FieldLevel) bool {
	return reflect.TypeOf(fl.Field().Interface()).Kind() == reflect.Bool
}

func (q fieldError) String() string {
	var sb strings.Builder

	sb.WriteString("Validation failed on field '" + q.err.Field() + "'")
	sb.WriteString(", condition: " + q.err.ActualTag())

	if q.err.Param() != "" {
		sb.WriteString(" { " + q.err.Param() + " }")
	}

	if q.err.Value() != nil && q.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", q.err.Value()))
	}

	return sb.String()
}

// function convert string to time.Time
func ConvertStringToTime(date string) (time.Time, error) {
	var response error
	result, err := time.Parse(time.RFC3339, date)
	if err != nil {
		response = errors.New("Date is wrong format")
		return result, response
	}

	return result, response
}

// validate date range by start date and end date
func ValidateDateRange(startDate time.Time, endDate time.Time) error {
	var response error

	// end date can equal to start date but not before
	if endDate.Before(startDate) {
		response = errors.New("End date must be after start date")
		return response
	}

	return response
}

// Custom validation functions
func validateThaiOrEnglish(fl validator.FieldLevel) bool {
	for _, char := range fl.Field().String() {
		if unicode.IsDigit(char) {
			return false
		}
		if !(unicode.Is(unicode.Thai, char) || unicode.IsLetter(char) || unicode.IsSpace(char)) {
			return false
		}
	}
	return true
}


func validateASCII(fl validator.FieldLevel) bool {
	for _, char := range fl.Field().String() {
		if char > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func validateComplexPassword(fl validator.FieldLevel) bool {
	var (
		hasMinLen  = false
		hasMaxLen  = true
		hasNumber  = false
		hasUpper   = false
		hasLower   = false
		hasSpecial = false
	)
	pass := fl.Field().String()

	if len(pass) >= 8 {
		hasMinLen = true
	}

	if len(pass) > 32 {
		hasMaxLen = false
	}

	for _, char := range pass {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case strings.ContainsRune("!@#$%^&*()-_+=", char):
			hasSpecial = true
		}
	}

	return hasMinLen && hasMaxLen && hasNumber && hasUpper && hasLower && hasSpecial
}
