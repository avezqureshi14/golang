package main

import app "payment/internal/app"


func main() {
	app := app.NewApp()
	app.Run()
}