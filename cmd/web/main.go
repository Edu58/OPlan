package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/Edu58/Oplan/config"
	"github.com/Edu58/Oplan/internal/app"
	"github.com/Edu58/Oplan/pkg/logger"
)

func main() {
	seedDB := flag.Bool("seed", false, "Run database seeds")
	seedType := flag.String("seed-type", "all", "Type of seed to run: (users, events, event_types)")

	flag.Parse()

	var loggerPath *os.File
	appConfig, err := config.LoadConfig(".", "app", "env")

	if err != nil {
		log.Fatalf("Could not load config with err: %v", err)
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
		}

		loggerPath = logFile
	}

	customLogger := logger.NewLoggerWithLevel(appConfig.LoggerLevel, loggerPath)

	myApp, err := app.NewApp(&appConfig, customLogger)

	if err != nil {
		log.Fatalf("Could create app with err: %v", err)
	}

	if *seedDB {
		// Use background context (no timeout) for seeding
		if err := myApp.RunSeeds(context.Background(), *seedType); err != nil {
			customLogger.WithField("MSG", "error running database seeds").Err(err)
			return
		}

		customLogger.Info("Database seeded successfuly")
		return
	}

	myApp.RunHTTP()
}
