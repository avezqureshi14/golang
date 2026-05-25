package todo

import (
	models "go-todo-app/internal/todo/models"
	repository "go-todo-app/internal/todo/repository"
	appErr "go-todo-app/pkg/errors"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetAll() ([]models.Todo, error) {
	todos, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoService) Get(id int) (models.Todo, error) {
	todo, err := s.repo.Get(id)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}

func (s *TodoService) Create(title string) (models.Todo, error) {
	if title == "" {
		return models.Todo{}, appErr.ErrInvalidInput
	}

	todo, err := s.repo.Create(title)

	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) Update(id int, title string, done bool) (models.Todo, error) {
	if title == "" {
		return models.Todo{}, appErr.ErrInvalidInput
	}

	todo, err := s.repo.Update(id, title, done)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
