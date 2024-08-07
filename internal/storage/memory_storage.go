package storage

import (
	"applicationDesign/internal/models"
	"context"

	"github.com/rs/zerolog"
)

type MemoryStorage struct {
	lg zerolog.Logger
}

var _ Storage = &MemoryStorage{}

func newMemoryStorage(lg zerolog.Logger) (Storage, error) {
	storage := &MemoryStorage{
		lg: lg.With().Caller().Logger(),
	}

	storage.lg.Info().Msg("Memory storage created")

	return storage, nil
}

func (s *MemoryStorage) Ping() error {
	return nil
}

func (s *MemoryStorage) Orders(ctx context.Context, order *models.Order) error {
	// booking
	return nil
}
