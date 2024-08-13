package provider

import (
	"applicationDesign/internal/config"
	"applicationDesign/internal/logic"
	"applicationDesign/internal/logic/rental"
	"applicationDesign/internal/logic/rental/accommodation"
	"applicationDesign/internal/logic/rental/rental_manager"
	"applicationDesign/internal/models"
	"context"

	"github.com/rs/zerolog"
)

type MemoryProvider struct {
	// @todo workaround, good way to use booking manager
	// for adding hotel.
	rentalManager rental_manager.BaseRentalManager

	bookingManager logic.BaseBookingManager
	lg             zerolog.Logger
	cfg            config.ServiceConfig
}

var _ Provider = &MemoryProvider{}

func newMemoryProvider(rentalManager rental_manager.BaseRentalManager, lg zerolog.Logger, cfg config.ServiceConfig) (Provider, error) {
	storage := &MemoryProvider{
		rentalManager:  rentalManager,
		bookingManager: logic.NewBookingManager(rentalManager, cfg.Workers, lg),
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

func (s *MemoryProvider) AddHotel(ctx context.Context, hotel *models.AddHotel) error {
	s.lg.Info().Msg("MemoryProvider: call AddHotel")

	h := rental.NewHotel(rental.TRentalID(hotel.HotelID), s.lg)
	for _, roomId := range hotel.RoomsID {
		h.AddAccommodation(accommodation.TAccommodationID(roomId))
	}

	s.rentalManager.AddRental(h)

	return nil
}
