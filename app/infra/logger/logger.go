package logger

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
)

func Debug(msg string) {
	logger := slog.New(createJsonHandler())
	_, file, line, _ := runtime.Caller(1)
	logger.Debug(msg, "path", fmt.Sprintf("%s:%d", file, line))
}

func Info(msg string) {
	logger := slog.New(createJsonHandler())
	_, file, line, _ := runtime.Caller(1)
	logger.Info(msg, "path", fmt.Sprintf("%s:%d", file, line))
}

func Error(msg string) {
	logger := slog.New(createJsonHandler())
	_, file, line, _ := runtime.Caller(1)
	logger.Error(msg, "path", fmt.Sprintf("%s:%d", file, line))
}

func createJsonHandler() *slog.JSONHandler {
	ops := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	h := slog.NewJSONHandler(os.Stdout, &ops)

	return h
}