package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Edu58/Oplan/config"
	db "github.com/Edu58/Oplan/internal/database/sqlc"
	httphandlers "github.com/Edu58/Oplan/internal/http_handlers"
)

type App struct {
	config  *config.Config
	queries *db.Queries
	pgxPool db.DBTX
	server  *http.Server
	mux     *http.ServeMux
}

func NewApp(config *config.Config, pgxPool db.DBTX) (*App, error) {
	mux := http.NewServeMux()
	addr := config.HOST + ":" + config.PORT

	return &App{
		config:  config,
		pgxPool: pgxPool,
		queries: db.New(pgxPool),
		mux:     mux,
		server: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}, nil
}

func (app *App) Init() error {
	log.Println("Setting up routes")
	httphandlers.NewDefaultHandler().RegisterRoutes(app.mux)
	return nil
}

func (app *App) Start() error {
	log.Printf("Starting server on %s", app.server.Addr)
	return app.server.ListenAndServe()
}

func (app *App) Shutdown(ctx context.Context, waitForShutdownCompletion chan struct{}) {
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigch

	log.Printf("Got signal: %v . Server shutting down.", sig)

	if err := app.server.Shutdown(ctx); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
	waitForShutdownCompletion <- struct{}{}
}
