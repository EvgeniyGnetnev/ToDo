package service

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/EvgeniyGnetnev/ToDo"
	"github.com/EvgeniyGnetnev/ToDo/pkg/repository"
)

// набор случайных символов для пароля (соль)
const salt = "asefiouyhg12q3214x"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
