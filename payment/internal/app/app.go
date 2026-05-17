package payment

import (
	"fmt"
	"payment/internal/platform/db"
	metrics "payment/internal/platform/metrics"
	profiler "payment/internal/platform/profiler"
	"payment/pkg/configs"
	"payment/pkg/logger"
	middleware "payment/pkg/metrics"

	payment "payment/internal/payment"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type App struct {
	router *gin.Engine
}

func NewApp() *App {
	r := gin.New()

	// Recovery middleware
	r.Use(gin.Recovery())

	// Profiler (pprof)
	profiler.Register(r)

	// Logger
	logger.Init()
	r.Use(logger.LoggerMiddleware(logger.Log))

	// Metrics
	metrics.Init()
	r.Use(middleware.MetricsMiddleware())

	// Prometheus endpoint (IMPORTANT: no middleware on this ideally)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Config
	cfgs := configs.Load()

	// DB
	dbConn := db.NewDB(cfgs.DB_URL)
	db.RunMigrations(dbConn)

	// Modules
	paymentModule := payment.New(dbConn)
	paymentModule.RegisterRoutes(r)

	return &App{
		router: r,
	}
}

func (a *App) Run() {
	port := configs.Load().PORT

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Println("Server running on", addr)

	if err := a.router.Run(addr); err != nil {
		panic(err)
	}
}
