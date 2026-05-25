package app

import (
	"fmt"
	"go-todo-app/internal/auth"
	"go-todo-app/internal/platform/db"
	"go-todo-app/internal/todo"
	"go-todo-app/pkg/configs"
	"go-todo-app/pkg/logger"

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
	// Logger
	// =========================
	logger.Init()
	defer logger.Sync()
	r.Use(logger.LoggerMiddleware(logger.Log))

	// =========================
	// CONFIG
	// =========================
	cfg := configs.Load()

	// =========================
	// INFRA
	// =========================
	db.RunMigrations(cfg.DB_URL)
	dbConn := db.NewDB(cfg.DB_URL)

	// =========================
	// MODULES
	// =========================
	authModule := auth.New(dbConn,logger.Log)
	todoModule := todo.New(dbConn,logger.Log)

	// =========================
	// ROUTES
	// =========================
	authModule.RegisterRoutes(r)
	todoModule.RegisterRoutes(r)

	return &App{
		router: r,
		cfg:    cfg,
	}
}

func (a *App) Run() {
	a.router.Run(fmt.Sprintf(":%s", a.cfg.PORT))
}
