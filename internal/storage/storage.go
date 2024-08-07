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

func newSyncMemoryStorage(lg zerolog.Logger) (Storage, error) {
	impl, err := newMemoryStorage(lg)
	if err != nil {
		return nil, err
	}

	return impl, nil
}

func NewStorage(lg zerolog.Logger, cfg config.ServiceConfig) (Storage, error) {
	if cfg.IsMemoryStorage() {
		lg.Info().Msg("create db storage")
		return newSyncMemoryStorage(lg)
	}

	return nil, fmt.Errorf("doesn't support this storage type: %d", cfg.StorageType)
}
