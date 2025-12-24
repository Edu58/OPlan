package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Edu58/Oplan/config"
	"github.com/Edu58/Oplan/internal/database"
	"github.com/Edu58/Oplan/internal/domain"
	httphandlers "github.com/Edu58/Oplan/internal/http_handlers"
	"github.com/Edu58/Oplan/internal/repository"
	"github.com/Edu58/Oplan/internal/service"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AppInterface interface {
	InitApp() error
	InitDB()
	InitHandlers() error
	InitRepositories() error
	InitServices() error
	Shutdown(ctx context.Context, waitForShutdownCompletion chan struct{})
	Start() error
}

type App struct {
	config             *config.Config
	pgxPool            *pgxpool.Pool
	server             *http.Server
	mux                *http.ServeMux
	accountTypeRepo    domain.AccountTypeRepository
	accountTypeService domain.AccountTypeService
	logger             logger.Logger
}

func NewApp(config *config.Config, logger logger.Logger) (AppInterface, error) {
	mux := http.NewServeMux()
	addr := config.HOST + ":" + config.PORT

	return &App{
		mux:    mux,
		config: config,
		logger: logger,
		server: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}, nil
}

func (app *App) InitApp() error {
	app.InitDB()
	app.InitRepositories()
	app.InitServices()
	app.InitHandlers()

	return nil
}

func (app *App) InitDB() {
	pgxPool, err := database.InitDB(context.Background(), app.config, app.logger)

	if err != nil {
		app.logger.Err(err)
		return
	}

	app.pgxPool = pgxPool
}

func (app *App) InitRepositories() error {
	app.logger.Info("Setting up repositories")
	app.accountTypeRepo = repository.NewAccountTypeRepository(app.pgxPool)
	return nil
}

func (app *App) InitServices() error {
	app.logger.Info("Setting up services")
	app.accountTypeService = service.NewAccountTypesService(app.accountTypeRepo, app.logger)
	return nil
}

func (app *App) InitHandlers() error {
	app.logger.Info("Setting up http handlers")
	accountTypeHandler := httphandlers.NewAccountTypesHandler(app.accountTypeService)
	accountTypeHandler.RegisterRoutes(app.mux)
	return nil
}

func (app *App) Start() error {
	app.logger.WithField("HOST", app.config.HOST).
		WithField("PORT", app.config.PORT).
		Info("Starting server")

	return app.server.ListenAndServe()
}

func (app *App) Shutdown(ctx context.Context, waitForShutdownCompletion chan struct{}) {
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	_ = <-sigch

	app.logger.Warn("Received shutdown signal")

	if err := app.server.Shutdown(ctx); err != nil {
		app.logger.Err(err)
	}
	waitForShutdownCompletion <- struct{}{}
}
