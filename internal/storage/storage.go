package storage

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/logic/guest_house"
	"applicationDesign/internal/models"
	"context"
	"fmt"

	"github.com/rs/zerolog"
)

type Storage interface {
	Ping() error
	Orders(ctx context.Context, order *models.Order) error
}

func NewStorage(guestHouseManager guest_house.GuestHouseManager, cfg config.ServiceConfig, lg zerolog.Logger) (Storage, error) {
	if cfg.IsMemoryStorage() {
		lg.Info().Msg("create db storage")
		return newMemoryStorage(guestHouseManager, lg, cfg)
	}

	return nil, fmt.Errorf("doesn't support this storage type: %d", cfg.StorageType)
}
