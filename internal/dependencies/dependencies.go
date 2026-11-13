package dependencies

import (
	"log/slog"
	"streamgate/internal/config"
	"streamgate/internal/httpclient"
)

type Base interface {
	Logger() *slog.Logger
}

type Service interface {
	Base
	Config() *config.Config
	HTTPClient() *httpclient.Client
}
