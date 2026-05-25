package todo

import (
	dto "go-todo-app/internal/todo/dto"
	models "go-todo-app/internal/todo/models"
)

func ToResponse(t models.Todo) dto.TodoResponse {
	return dto.TodoResponse{
		ID:    t.ID,
		Title: t.Title,
		Done:  t.Done,
	}
}