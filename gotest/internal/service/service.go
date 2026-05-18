package service_test

import (
	"errors"
	"gotest/internal/model"
	"gotest/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Register(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	_, err := s.repo.GetByName(name)
	if err == nil {
		return errors.New("user already exists")
	}

	return s.repo.Save(model.User{Name: name})
}
