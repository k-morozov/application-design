package storage

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/logic"
	"applicationDesign/internal/logic/guest_house"
	"applicationDesign/internal/models"
	"context"

	"github.com/rs/zerolog"
)

type MemoryStorage struct {
	bookingManager logic.Manager
	lg             zerolog.Logger
	cfg            config.ServiceConfig
}

var _ Storage = &MemoryStorage{}

func newMemoryStorage(guestHouseManager guest_house.GuestHouseManager, lg zerolog.Logger, cfg config.ServiceConfig) (Storage, error) {
	storage := &MemoryStorage{
		bookingManager: logic.NewBookingManager(guestHouseManager, cfg.Workers, lg),
		lg:             lg.With().Caller().Logger(),
		cfg:            cfg,
	}

	storage.lg.Info().Msg("Memory storage created")

	return storage, nil
}

func (s *MemoryStorage) Ping() error {
	return nil
}

func (s *MemoryStorage) Orders(ctx context.Context, order *models.Order) error {
	var bookingId logic.BookingID
	var err error

	s.lg.Info().Msg("MemoryStorage: call Orders")

	if bookingId, err = s.bookingManager.PrepareBook(order); err != nil {
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
