package payment

import (
	"fmt"
	"payment/internal/platform/db"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func NewApp()* App{
	r := gin.New()
	r.Use(gin.Recovery())

	dbConn := db.NewDB("")

	return &App{
		router: r,
	}
}

func (a *App) Run() {
	a.router.Run(fmt.Sprintf("localhost:%s","8000"))
}