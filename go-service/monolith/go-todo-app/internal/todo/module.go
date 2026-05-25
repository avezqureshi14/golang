package todo

import (
	todo_handler "go-todo-app/internal/todo/handler"
	todo_repo "go-todo-app/internal/todo/repository"
	todo_routes "go-todo-app/internal/todo/routes"
	todo_service "go-todo-app/internal/todo/service"
	"go-todo-app/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Module struct {
	handler *todo_handler.TodoHandler
	mw      gin.HandlerFunc

}

func New(db *sqlx.DB, log *zap.Logger) *Module {
	repo := todo_repo.NewTodoRepo(db,log)
	service := todo_service.NewTodoService(repo)
	handler := todo_handler.NewTodoHandler(service)

	mw := jwt.JWT()

	return &Module{
		handler: handler,
		mw:      mw,
	}
}

func (m *Module) RegisterRoutes(r *gin.Engine) {
	todo_routes.RegisterTodoRoutes(r, m.handler, m.mw)
}