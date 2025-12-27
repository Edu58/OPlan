package logger

import (
	"io"
	"strings"

	"github.com/rs/zerolog"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Err(err error)
	Error(msg string)
	Fatal(msg string)
	GetLevel() string
	WithField(key string, value any) Logger
	WithFields(fields map[string]any) Logger
}

type ZeroLogger struct {
	zerolog.Logger
}

func NewLogger(out io.Writer) Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	return &ZeroLogger{zerolog.
		New(out).
		With().
		Timestamp().
		Logger(),
	}
}

func NewLoggerWithLevel(level string, out io.Writer) Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Set the global log level based on the provided level
	switch strings.ToLower(level) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn", "warning":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "disabled", "off":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	default:
		// Default to info level if unknown level is provided
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	return &ZeroLogger{
		zerolog.
			New(out).
			Level(zerolog.GlobalLevel()).
			With().
			Timestamp().
			Logger(),
	}
}

func (l *ZeroLogger) GetLevel() string {
	return l.Logger.GetLevel().String()
}

func (l *ZeroLogger) Info(msg string) {
	l.Logger.Info().Msg(msg)
}

func (l *ZeroLogger) Debug(msg string) {
	l.Logger.Debug().Msg(msg)
}

func (l *ZeroLogger) Warn(msg string) {
	l.Logger.Warn().Msg(msg)
}

func (l *ZeroLogger) Err(err error) {
	l.Logger.Err(err).Msg(err.Error())
}

func (l *ZeroLogger) Error(msg string) {
	l.Logger.Error().Msg(msg)
}

func (l *ZeroLogger) Fatal(msg string) {
	l.Logger.Fatal().Msg(msg)
}

func (l *ZeroLogger) WithField(key string, value any) Logger {
	return &ZeroLogger{Logger: l.Logger.With().Interface(key, value).Logger()}
}

func (l *ZeroLogger) WithFields(fields map[string]any) Logger {
	ctx := l.Logger.With()

	for k, v := range fields {
		ctx.Interface(k, v)
	}

	newLogger := ctx.Logger()

	return &ZeroLogger{newLogger}
}
