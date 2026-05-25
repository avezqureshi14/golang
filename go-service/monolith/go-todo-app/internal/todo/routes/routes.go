package todo

import (
	handler "go-todo-app/internal/todo/handler"

	"github.com/gin-gonic/gin"
)


func RegisterTodoRoutes(r *gin.Engine, h *handler.TodoHandler, jwt gin.HandlerFunc) {
	todo := r.Group("/todos")
	todo.Use(jwt)

	todo.GET("", h.GetTodos)
	todo.POST("", h.CreateTodo)
	todo.PUT("/:id", h.UpdateTodo)
	todo.DELETE("/:id", h.DeleteTodo)
}