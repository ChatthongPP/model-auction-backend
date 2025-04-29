package repositories

import (
	"errors"

	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

func (r *Repo) CreateUser(req domain.User) (*domain.User, error) {
	user := models.User{
		ID:          req.ID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		Gender:      req.Gender,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		CitizenID:   req.CitizenID,
		RoleID:      req.RoleID,
	}

	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	res := domain.MapUserToDomain(user)

	return &res, nil
}

func (r *Repo) FindByEmail(email string) (*domain.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	res := domain.MapUserToDomain(user)

	return &res, nil
}

func (r *Repo) GetUserByID(userID uint) (*domain.User, error) {
	var user models.User
	err := r.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	res := domain.MapUserToDomain(user)
	return &res, nil
}
