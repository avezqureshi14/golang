package payment

import (
	"fmt"
	"payment/internal/platform/db"
	"payment/pkg/configs"
	"payment/pkg/logger"

	payment "payment/internal/payment"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func NewApp() *App {
	r := gin.New()
	r.Use(gin.Recovery())

	logger.Init()
	defer logger.Sync()
	r.Use(logger.LoggerMiddleware(logger.Log))

	cfgs := configs.Load()

	dbConn := db.NewDB(cfgs.DB_URL)
	db.RunMigrations(dbConn)

	paymentModule := payment.New(dbConn)

	paymentModule.RegisterRoutes(r)

	return &App{
		router: r,
	}
}

func (a *App) Run() {
	a.router.Run(fmt.Sprintf("localhost:%s", configs.Load().PORT))
}
