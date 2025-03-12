package service

import (
	"go.challenge/repository"
	"go.challenge/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(username string, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)
	if (err != nil) {
		return "", err
	}
	if (user == nil) {
		return "", nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}
	jwt, err := utils.GenerateJWT(username)
	if err != nil {
		return "", err
	}
	return jwt, nil
}
