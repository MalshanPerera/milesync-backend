package utils

import (
	"os"

	"github.com/rs/zerolog"
)

type logger interface {
	Info()
	Error()
}

type Logger struct {
	logger zerolog.Logger
}

var appLogger Logger

func (l *Logger) Info(msg string) {
	appLogger.logger.Info().Msg(msg)
}

func (l *Logger) Error(msg string) {
	appLogger.logger.Error().Msg(msg)
}

func InitLogger() Logger {
	zeroLogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	appLogger = Logger{
		logger: zeroLogger,
	}

	return GetLogger()
}

func GetLogger() Logger {
	return appLogger
}

func GetZeroLogger() zerolog.Logger {
	return appLogger.logger
}
