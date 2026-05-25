package auth

import (
	auth_handler "go-todo-app/internal/auth/handler"
	auth_repo "go-todo-app/internal/auth/repository"
	auth_routes "go-todo-app/internal/auth/routes"
	auth_service "go-todo-app/internal/auth/service"
	"go-todo-app/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Module struct {
	handler *auth_handler.AuthHandler
	mw      gin.HandlerFunc
}

func New(db *sqlx.DB, log *zap.Logger) *Module {
	repo := auth_repo.NewAuthRepo(db)
	service := auth_service.NewAuthService(repo)
	handler := auth_handler.NewAuthHandler(service)

	mw := jwt.JWT()

	return &Module{
		handler: handler,
		mw:      mw,
	}
}

func (m *Module) RegisterRoutes(r *gin.Engine) {
	auth_routes.RegisterAuthRoutes(r, m.handler, m.mw)
}