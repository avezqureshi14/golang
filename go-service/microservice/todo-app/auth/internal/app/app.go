package app

import (
	"fmt"
	"go-todo-app/internal/auth"
	"go-todo-app/internal/platform/db"
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
	authModule := auth.New(dbConn)

	// =========================
	// ROUTES
	// =========================
	authModule.RegisterRoutes(r)

	return &App{
		router: r,
		cfg:    cfg,
	}
}

func (a *App) Run() {
	a.router.Run(fmt.Sprintf(":%s", a.cfg.PORT))
}