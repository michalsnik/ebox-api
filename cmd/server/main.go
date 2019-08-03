package main

import (
	"ebox-api/internal/app"
	"ebox-api/internal/config"
)

func main() {
	cfg := config.GetConfig()
	server := app.Create(cfg)
	server.Run(":4000")
}
