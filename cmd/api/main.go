package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Edu58/Oplan/config"
	"github.com/Edu58/Oplan/internal/database"
)

func main() {
	appConfig, err := config.LoadConfig(".", "app", "env")

	if err != nil {
		log.Fatalf("Could not load config with err: %v", err)
		return
	}

	_, err = database.InitDB(context.Background(), &appConfig)

	// store := db.New(dbPoolConn)

	if err != nil {
		log.Fatalf("Could not load config with err: %v", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(appConfig.HOST+":"+appConfig.PORT, nil))
}
