package storage

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/models"
	"context"
	"fmt"

	"github.com/rs/zerolog"
)

type Storage interface {
	Ping() error
	Orders(ctx context.Context, order *models.Order) error
}

func NewStorage(lg zerolog.Logger, cfg config.ServiceConfig) (Storage, error) {
	if cfg.IsMemoryStorage() {
		lg.Info().Msg("create db storage")
		return newMemoryStorage(lg, cfg)
	}

	return nil, fmt.Errorf("doesn't support this storage type: %d", cfg.StorageType)
}
