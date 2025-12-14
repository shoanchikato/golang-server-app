package service

import (
	"log/slog"
)

type Logger interface {
	Error(msg string, args ...any)
}

type log struct {
	logger *slog.Logger
}

func NewLogger(logger *slog.Logger) Logger {
	return &log{logger}
}

func (l *log) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
