package log

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

const logTimeFormat = "2006-01-02T15:04:05.999"

type loggerKey struct{}

func NewLogger(level string) zerolog.Logger {
	loggerLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		loggerLevel = zerolog.InfoLevel
	}
	multi := zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: logTimeFormat},
		//zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: logTimeFormat},
	)

	return zerolog.New(multi).Level(loggerLevel).With().Timestamp().Logger()
}

func DefaultLogger() zerolog.Logger {
	return NewLogger("info")
}

func UpdateContext(ctx context.Context, logger zerolog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func FromContext(ctx context.Context) zerolog.Logger {
	lg, ok := ctx.Value(loggerKey{}).(zerolog.Logger)
	if !ok {
		return DefaultLogger()
	}
	return lg
}
