package provider

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/logic/hotel/manager"
	"applicationDesign/internal/models"
	"context"
	"fmt"

	"github.com/rs/zerolog"
)

type Provider interface {
	Ping() error
	Orders(ctx context.Context, order *models.Order) error
}

func NewProvider(guestHouseManager manager.BaseHotelManager, cfg config.ServiceConfig, lg zerolog.Logger) (Provider, error) {
	if cfg.IsMemoryStorage() {
		lg.Info().Msg("create db provider")
		return newMemoryProvider(guestHouseManager, lg, cfg)
	}

	return nil, fmt.Errorf("doesn't support this provider type: %d", cfg.StorageType)
}
