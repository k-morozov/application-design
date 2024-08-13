package provider

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/logic/rental/rental_manager"
	"applicationDesign/internal/models"
	"context"
	"fmt"

	"github.com/rs/zerolog"
)

type Provider interface {
	Ping() error
	Orders(ctx context.Context, order *models.Order) error
	AddHotel(ctx context.Context, hotel *models.AddHotel) error
}

func NewProvider(guestHouseManager rental_manager.BaseRentalManager, cfg config.ServiceConfig, lg zerolog.Logger) (Provider, error) {
	if cfg.IsMemoryStorage() {
		lg.Info().Msg("create db provider")
		return newMemoryProvider(guestHouseManager, lg, cfg)
	}

	return nil, fmt.Errorf("doesn't support this provider type: %d", cfg.StorageType)
}
