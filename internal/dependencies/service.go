package dependencies

import (
	"log/slog"
	"streamgate/internal/config"
	"streamgate/internal/httpclient"
	"time"
)

type service struct {
	Base
	config     *config.Config
	httpClient *httpclient.Client
}

func New(cfg *config.Config, logger *slog.Logger) Service {
	return &service{
		Base:       NewBase(logger),
		config:     cfg,
		httpClient: httpclient.New(httpclient.WithTimeout(time.Second * 5)),
	}
}

func (s *service) Config() *config.Config {
	return s.config
}

func (s *service) HTTPClient() *httpclient.Client {
	return s.httpClient
}
