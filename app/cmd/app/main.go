package main

import (
	"log"

	"github.com/goocarry/bootstrapper/internal/app"
	"github.com/goocarry/bootstrapper/internal/config"
	"github.com/goocarry/bootstrapper/pkg/logger"
)

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()

	log.Print("logger initializing")
	logger := logger.GetLogger(cfg.AppConfig.LogLevel)

	app, err := app.NewApp(cfg, &logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Running Application")
	app.Run()
}
