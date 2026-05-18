package service

import (
	"errors"
	"gotest/internal/model"
	service "gotest/internal/service"

	"testing"
)

type MockRepo struct {
	users map[string]model.User
}

func NewMockRepo() *MockRepo {
	return &MockRepo{
		users: make(map[string]model.User),
	}
}

func (m *MockRepo) Save(user model.User) error {
	m.users[user.Name] = user
	return nil
}

func (m *MockRepo) GetByName(name string) (model.User, error) {
	user, ok := m.users[name]
	if !ok {
		return model.User{}, errors.New("not found")
	}
	return user, nil
}

func TestRegister(t *testing.T) {
	repo := NewMockRepo()
	svc := service.NewUserService(repo)

	// success case
	err := svc.Register("avez")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// duplicate case
	err = svc.Register("avez")
	if err == nil {
		t.Errorf("expected duplicate error")
	}

	// empty name
	err = svc.Register("")
	if err == nil {
		t.Errorf("expected empty name error")
	}
}
