package dependencies

import (
	"log/slog"
)

type base struct {
	logger *slog.Logger
}

func NewBase(logger *slog.Logger) Base {
	return &base{
		logger: logger,
	}
}

func (b *base) Logger() *slog.Logger {
	return b.logger
}
