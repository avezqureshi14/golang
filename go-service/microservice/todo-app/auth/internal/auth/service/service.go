package service

import (
	"errors"
	dto "go-todo-app/internal/auth/dto"
	models "go-todo-app/internal/auth/models"
	repository "go-todo-app/internal/auth/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) Signup(req dto.SignupRequest) (models.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}

	return s.repo.Create(user)
}

func (s *AuthService) Signin(req dto.SigninRequest) (models.User, error) {
	user, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return user, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *AuthService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}