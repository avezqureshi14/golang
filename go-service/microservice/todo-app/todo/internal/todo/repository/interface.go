package todo

import (
	models "go-todo-app/internal/todo/models"
)
type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	Get(id int) (models.Todo, error)
	Create(title string) (models.Todo, error)
	Update(id int, title string, done bool) (models.Todo, error)
	Delete(id int) error
}
