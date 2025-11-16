package main

import (
	"context"
	"log"
	"time"

	"github.com/Edu58/Oplan/config"
	"github.com/Edu58/Oplan/internal/app"
	"github.com/Edu58/Oplan/internal/database"
)

func main() {
	appConfig, err := config.LoadConfig(".", "app", "env")

	if err != nil {
		log.Fatalf("Could not load config with err: %v", err)
		return
	}

	dbPool, err := database.InitDB(context.Background(), &appConfig)

	app, err := app.NewApp(&appConfig, dbPool)

	if err != nil {
		log.Fatalf("Could create app with err: %v", err)
	}

	if err := app.Init(); err != nil {
		log.Fatalf("Error initializing app: %v", err)
	}
	waitForShutdownCompletion := make(chan struct{})
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	go app.Shutdown(ctx, waitForShutdownCompletion)
	defer cancel()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	<-waitForShutdownCompletion
}
