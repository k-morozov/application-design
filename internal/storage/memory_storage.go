package storage

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/logic"
	"applicationDesign/internal/models"
	"context"

	"github.com/rs/zerolog"
)

type MemoryStorage struct {
	manager logic.Manager
	lg      zerolog.Logger
	cfg     config.ServiceConfig
}

var _ Storage = &MemoryStorage{}

func newMemoryStorage(lg zerolog.Logger, cfg config.ServiceConfig) (Storage, error) {
	storage := &MemoryStorage{
		manager: logic.NewBookingManager(lg, cfg.Workers),
		lg:      lg.With().Caller().Logger(),
		cfg:     cfg,
	}

	storage.lg.Info().Msg("Memory storage created")

	return storage, nil
}

func (s *MemoryStorage) Ping() error {
	return nil
}

func (s *MemoryStorage) Orders(ctx context.Context, order *models.Order) error {
	var booking_id logic.BookingID
	var err error

	s.lg.Info().Msg("MemoryStorage: call Orders")

	if booking_id, err = s.manager.PrepareBook(order); err != nil {
		s.lg.Error().Msg("failed prepared booking")
		return err
	}

	s.lg.Info().Str("booking_id", booking_id.String()).Msg("successfully prepared")

	if err = s.manager.AcceptBook(booking_id); err != nil {
		s.lg.Error().Str("booking_id", booking_id.String()).Msg("failed booked")
		return err
	}

	s.lg.Info().Str("booking_id", booking_id.String()).Msg("successfully booked")
	return nil
}
