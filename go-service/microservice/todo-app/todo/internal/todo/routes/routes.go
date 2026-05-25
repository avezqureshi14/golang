package todo

import (
	handler "go-todo-app/internal/todo/handler"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(r *gin.Engine, h *handler.TodoHandler, jwt gin.HandlerFunc) {
	r.GET("/", h.GetTodos)
	r.POST("/", h.CreateTodo)
	r.PUT("/:id", h.UpdateTodo)
	r.DELETE("/:id", h.DeleteTodo)
}
