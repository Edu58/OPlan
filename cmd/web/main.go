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
	var loggerPath *os.File
	appConfig, err := config.LoadConfig(".", "app", "env")

	if err != nil {
		log.Fatalf("Could not load config with err: %v", err)
		return
	}

	if appConfig.LoggerPath == "" {
		loggerPath = os.Stdout
	} else {
		logFile, err := os.OpenFile(appConfig.LoggerPath,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0o644,
		)

		if err != nil {
			log.Fatalf("Could not create log file: %v", err)
			return
		}

		loggerPath = logFile
	}

	customLogger := logger.NewLoggerWithLevel(appConfig.LoggerLevel, loggerPath)

	myApp, err := app.NewApp(&appConfig, customLogger)

	if err != nil {
		log.Fatalf("Could create app with err: %v", err)
	}

	if err := myApp.InitApp(); err != nil {
		log.Fatalf("Error initializing app: %v", err)
	}

	waitForShutdownCompletion := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	// Graceful Shutdown
	go myApp.Shutdown(ctx, waitForShutdownCompletion)
	defer cancel()

	if err := myApp.Start(); err != nil {
		log.Fatal(err)
	}

	<-waitForShutdownCompletion
}
