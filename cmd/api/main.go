package main

import (
	"log"
	"net/http"

	"github.com/Edu58/Oplan/config"
	"github.com/Edu58/Oplan/internal/database"
)

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("Could not load config with err: %v", err)
		return
	}

	_, err = database.InitDB(&config)

	// store := db.New(dbPoolConn)

	if err != nil {
		log.Fatalf("Could not load config with err: %v", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(config.HOST+":"+config.PORT, nil))
}
