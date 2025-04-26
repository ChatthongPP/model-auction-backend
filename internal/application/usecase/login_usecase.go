package usecase

import (
	"backend-service/internal/domain"


	"golang.org/x/crypto/bcrypt"
)

// Function to hash password
func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// Function to check password hash
func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (uc *Usecase) CheckEmailAndPassword(admin *domain.Login) error {
	return uc.loginRepo.GetAdminByEmailAndPassword(admin)
}
