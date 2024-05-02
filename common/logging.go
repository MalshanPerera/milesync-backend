package common

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type MyLogger struct {
	zerolog.Logger
}

var Logger MyLogger

func init() {
	// create output configuration
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	// Format level: fatal, error, debug, info, warn
	output.FormatLevel = func(i interface{}) string {
		color, reset := getColorByLevel(fmt.Sprintf("%s", i))
		return color + strings.ToUpper(fmt.Sprintf("| %-6s|", i)) + reset
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	// format error
	output.FormatErrFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s: ", i)
	}

	zerolog := zerolog.New(output).With().Caller().Timestamp().Logger()
	Logger = MyLogger{zerolog}
}

func getColorByLevel(level string) (string, string) {
	switch level {
	case "info":
		return "\033[36m", "\033[0m"
	case "error":
		return "\033[31m", "\033[0m"
	case "debug":
		return "\033[34m", "\033[0m"
	case "warn":
		return "\033[33m", "\033[0m"
	case "fatal":
		return "\033[31m", "\033[0m"
	default:
		return "\033[36m", "\033[0m"
	}
}

func (l *MyLogger) LogInfo() *zerolog.Event {
	return l.Logger.Info()
}

func (l *MyLogger) LogError() *zerolog.Event {
	return l.Logger.Error()
}

func (l *MyLogger) LogDebug() *zerolog.Event {
	return l.Logger.Debug()
}

func (l *MyLogger) LogWarn() *zerolog.Event {
	return l.Logger.Warn()
}

func (l *MyLogger) LogFatal() *zerolog.Event {
	return l.Logger.Fatal()
}
