package main

import (
	gateway_proxy "gateway/internal/proxy"
	gateway_routes "gateway/internal/routes"
	"gateway/pkg/configs"
	"gateway/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// services
	userService := gateway_proxy.NewProxy("http://localhost:9001")
	todoService := gateway_proxy.NewProxy("http://localhost:9002")

	// middleware
	jwtMiddleware := jwt.JWT()

	// collect routes
	var allRoutes []gateway_routes.Route

	allRoutes = append(allRoutes, gateway_routes.AuthRoutes(userService)...)
	allRoutes = append(allRoutes, gateway_routes.UserRoutes(userService)...)
	allRoutes = append(allRoutes, gateway_routes.TodoRoutes(todoService)...)

	// register
	gateway_routes.Register(r, jwtMiddleware, allRoutes)

	// start server
	cfg := configs.Load()
	r.Run(":" + cfg.PORT)
}