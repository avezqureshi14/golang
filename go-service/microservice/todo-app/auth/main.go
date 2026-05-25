package main

import "go-todo-app/internal/app"

func main() {
	app := app.NewApp()
	app.Run()
}