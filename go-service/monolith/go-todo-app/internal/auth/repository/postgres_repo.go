package auth

import (
	models "go-todo-app/internal/auth/models"

	"github.com/jmoiron/sqlx"
)

type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) AuthRepository {
	return &authRepo{db: db}
}

func (r *authRepo) Create(user models.User) (models.User, error) {
	query := `INSERT INTO users (name, email, password)
	          VALUES ($1,$2,$3) RETURNING id`

	err := r.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
	return user, err
}

func (r *authRepo) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.Get(&user, `SELECT * FROM users WHERE email=$1`, email)
	return user, err
}

func (r *authRepo) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Select(&users, `SELECT id, name, email FROM users`)
	return users, err
}