package provider

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/logic"
	"applicationDesign/internal/logic/renter/manager"
	"applicationDesign/internal/models"
	"context"

	"github.com/rs/zerolog"
)

type MemoryProvider struct {
	bookingManager logic.BaseBookingManager
	lg             zerolog.Logger
	cfg            config.ServiceConfig
}

var _ Provider = &MemoryProvider{}

func newMemoryProvider(guestHouseManager manager.BaseRentersManager, lg zerolog.Logger, cfg config.ServiceConfig) (Provider, error) {
	storage := &MemoryProvider{
		bookingManager: logic.NewBookingManager(guestHouseManager, cfg.Workers, lg),
		lg:             lg.With().Caller().Logger(),
		cfg:            cfg,
	}

	storage.lg.Info().Msg("Memory provider created")

	return storage, nil
}

func (s *MemoryProvider) Ping() error {
	return nil
}

func (s *MemoryProvider) Orders(ctx context.Context, order *models.Order) error {
	var bookingId logic.TBookingID
	var err error

	s.lg.Info().Msg("MemoryProvider: call Orders")

	if bookingId, err = s.bookingManager.PrepareBook(*order); err != nil {
		s.lg.Error().Msg("failed prepared booking")
		return err
	}

	s.lg.Info().Str("booking_id", bookingId.String()).Msg("successfully prepared")

	if err = s.bookingManager.AcceptBook(bookingId); err != nil {
		s.lg.Error().Str("booking_id", bookingId.String()).Msg("failed booked")
		return err
	}

	s.lg.Info().Str("booking_id", bookingId.String()).Msg("successfully booked")
	return nil
}
