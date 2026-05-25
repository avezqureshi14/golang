package app

import (
	"fmt"
	"go-todo-app/internal/platform/db"
	"go-todo-app/internal/todo"
	"go-todo-app/pkg/configs"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
	cfg    configs.Config
}

func NewApp() *App {
	r := gin.New()
	r.Use(gin.Recovery())

	// =========================
	// CONFIG
	// =========================
	cfg := configs.Load()

	// =========================
	// INFRA
	// =========================
	dbConn := db.NewDB(cfg.DB_URL)

	// =========================
	// MODULES
	// =========================
	todoModule := todo.New(dbConn)

	// =========================
	// ROUTES
	// =========================
	todoModule.RegisterRoutes(r)

	return &App{
		router: r,
		cfg:    cfg,
	}
}

func (a *App) Run() {
	a.router.Run(fmt.Sprintf(":%s", a.cfg.PORT))
}