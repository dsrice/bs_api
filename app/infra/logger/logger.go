package logger

import (
	"log/slog"
	"os"
)

func Debug(msg string) {
	logger := slog.New(createJsonHandler())
	logger.Debug(msg)
}

func Info(msg string) {
	logger := slog.New(createJsonHandler())
	logger.Info(msg)
}

func Error(msg string) {
	logger := slog.New(createJsonHandler())
	logger.Error(msg)
}

func createJsonHandler() *slog.JSONHandler {
	h := slog.NewJSONHandler(os.Stdout, nil)

	return h
}