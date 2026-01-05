package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Edu58/Oplan/config"
	"github.com/Edu58/Oplan/internal/database"
	"github.com/Edu58/Oplan/internal/database/seeds"
	"github.com/Edu58/Oplan/internal/database/sqlc"
	httphandlers "github.com/Edu58/Oplan/internal/http_handlers"
	"github.com/Edu58/Oplan/internal/services"
	"github.com/Edu58/Oplan/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Pre-compile the regex (only once)
// var msisdnRegex = regexp.MustCompile(`^\+\d{1,3}\d{9,}$`)

type AppInterface interface {
	InitApp() error
	InitDB() error
	InitServices()
	InitHandlers()
	Shutdown(ctx context.Context, waitForShutdownCompletion chan struct{})
	Start() error
}

type App struct {
	config           *config.Config
	pgxPool          *pgxpool.Pool
	server           *http.Server
	mux              *http.ServeMux
	store            *sqlc.Queries
	sessionService   *services.SessionService
	userService      *services.UserService
	otpService       *services.OTPService
	eventTypeService *services.EventTypeService
	eventService     *services.EventService
	logger           logger.Logger
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
	err := app.InitDB()

	if err != nil {
		return err
	}

	app.InitServices()
	app.InitHandlers()

	return nil
}

func (app *App) InitDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pgxPool, err := database.InitDB(ctx, app.config, app.logger)
	if err != nil {
		app.logger.Err(err)
		return err
	}

	app.pgxPool = pgxPool
	app.store = sqlc.New(pgxPool)

	// Run seeds
	// Use background context (no timeout) for seeding
	if err := seeds.Run(context.Background(), app.store); err != nil {
		return err
	}

	return nil
}

func (app *App) InitServices() {
	app.sessionService = services.NewSessionService(app.store, app.logger)
	app.userService = services.NewUserService(app.store, app.logger)
	app.otpService = services.NewOTPHandler(app.store, app.logger)
	app.eventService = services.NewEventService(app.store, app.logger)
	app.eventTypeService = services.NewEventTypeService(app.store, app.logger)
}

func (app *App) InitHandlers() {
	app.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))

	indexHandler := httphandlers.NewIndexHandler(app.eventService, app.eventTypeService, app.logger)
	indexHandler.RegisterRoutes(app.mux)

	authHandler := httphandlers.NewSessionHandler(app.sessionService, app.userService, app.otpService, app.logger)
	authHandler.RegisterRoutes(app.mux)
}

func (app *App) Start() error {
	app.logger.WithField("HOST", app.config.HOST).
		WithField("PORT", app.config.PORT).
		Info("Starting server")

	return app.server.ListenAndServe()
}

func (app *App) Shutdown(ctx context.Context, waitForShutdownCompletion chan struct{}) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	_ = <-signalChan

	app.logger.Warn("Received shutdown signal")

	if err := app.server.Shutdown(ctx); err != nil {
		app.logger.Err(err)
	}
	waitForShutdownCompletion <- struct{}{}
}
