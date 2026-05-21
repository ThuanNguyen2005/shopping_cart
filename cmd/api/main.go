package main

import (
	"23.usersapi/internal/app"
	"23.usersapi/internal/config"
)

func main() {
	app.LoadEnv()
	//initial
	cfg := config.NewConfig()
	//initial application
	r := app.NewApplication(cfg)

	//start server
	if err := r.Run(); err != nil {
		panic("Failed to run server.")
	}
}
