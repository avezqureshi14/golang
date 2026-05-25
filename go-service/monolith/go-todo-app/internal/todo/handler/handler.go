package todo

import (
	http_response "go-todo-app/internal/platform/http"
	dto "go-todo-app/internal/todo/dto"
	service "go-todo-app/internal/todo/service"
	mapper "go-todo-app/internal/todo/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.service.GetAll()
	if err != nil {
		http_response.HandleError(c, err)
		return
	}
	http_response.SendSuccess(c, http.StatusOK, todos)
}

func (h *TodoHandler) GetTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_response.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}

	todo, err := h.service.Get(id)
	if err != nil {
		log.Print("Todo Handler ",err)
		http_response.HandleError(c, err)
		return
	}

	http_response.SendSuccess(c, http.StatusOK, mapper.ToResponse(todo))
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req dto.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http_response.SendError(c, http.StatusBadRequest, "invalid body")
		return
	}

	todo, err := h.service.Create(req.Title)
	if err != nil {
		http_response.HandleError(c, err)
		return
	}

	http_response.SendSuccess(c, http.StatusCreated, mapper.ToResponse(todo))
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_response.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}

	var req dto.UpdateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		http_response.SendError(c, http.StatusBadRequest, "invalid body")
		return
	}

	todo, err := h.service.Update(id, req.Title, req.Done)
	if err != nil {
		http_response.HandleError(c, err)
		return
	}

	http_response.SendSuccess(c, http.StatusOK, mapper.ToResponse(todo))
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http_response.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http_response.HandleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}