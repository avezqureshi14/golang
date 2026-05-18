// internal/repository/postgres_repo.go
package repository

import (
	"errors"
	"fmt"
	"gotest/internal/model"
)

type PostgresRepo struct{}

func NewPostgresRepo() *PostgresRepo {
	return &PostgresRepo{}
}

func (p *PostgresRepo) Save(user model.User) error {
	fmt.Println("Saving to DB:", user.Name)
	return nil
}

func (p *PostgresRepo) GetByName(name string) (model.User, error) {
	return model.User{}, errors.New("not found")
}
