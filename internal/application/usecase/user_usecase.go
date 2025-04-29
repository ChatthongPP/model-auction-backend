package usecase

import (
	"backend-service/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

func (uc *Usecase) CreateUser(req domain.User) (res *domain.User, err error) {
	// // Check if email already exists
	existingUser, err := uc.userRepo.FindByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, domain.ErrEmailAlreadyExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.Password = string(hashedPassword)

	if res, err = uc.userRepo.CreateUser(req); err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *Usecase) Login(email, password string) (*domain.User, error) {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return nil, domain.ErrEmailNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, domain.ErrIncorrectPassword
	}

	return user, nil
}

func (uc *Usecase) GetUserByID(userID uint) (*domain.User, error) {
	return uc.userRepo.GetUserByID(userID)
}
