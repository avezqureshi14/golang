package repository

import "gotest/internal/model"

type UserRepository interface {
	Save(user model.User) error
	GetByName(name string) (model.User, error)
}
