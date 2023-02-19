package main

import (
	"github.com/Mersock/project-timesheet-backend/config"
	"github.com/Mersock/project-timesheet-backend/internal/app"
	"log"
)

func main() {
	//config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	//Run server
	app.Run(cfg)
}
