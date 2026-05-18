// cmd/main.go
package main

import (
	"gotest/internal/repository"
	service "gotest/internal/service"
)

func main() {
	repo := repository.NewPostgresRepo()
	svc := service.NewUserService(repo)

	svc.Register("avez")
}
