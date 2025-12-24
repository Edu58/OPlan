package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Edu58/Oplan/config"
	"github.com/Edu58/Oplan/internal/app"
	"github.com/Edu58/Oplan/pkg/logger"
)

func main() {
	var logger_path *os.File
	appConfig, err := config.LoadConfig(".", "app", "env")

	if err != nil {
		log.Fatalf("Could not load config with err: %v", err)
		return
	}

	if appConfig.LOGGER_PATH == "" {
		logger_path = os.Stdout
	} else {
		log_file, err := os.OpenFile(appConfig.LOGGER_PATH,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0o644,
		)

		if err != nil {
			log.Fatalf("Could not create log file: %v", err)
			return
		}

		logger_path = log_file
	}

	logger := logger.NewLoggerWithLevel(appConfig.LOGGER_LEVEL, logger_path)

	app, err := app.NewApp(&appConfig, logger)

	if err != nil {
		log.Fatalf("Could create app with err: %v", err)
	}

	if err := app.InitApp(); err != nil {
		log.Fatalf("Error initializing app: %v", err)
	}

	waitForShutdownCompletion := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	// Graceful Shutdown
	go app.Shutdown(ctx, waitForShutdownCompletion)
	defer cancel()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	<-waitForShutdownCompletion
}
