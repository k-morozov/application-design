package logic

import (
	"applicationDesign/internal/models"

	"github.com/rs/zerolog"
)

type BookingManager struct {
	lg zerolog.Logger
}

var _ Manager = &BookingManager{}

func NewBookingManager(lg zerolog.Logger) Manager {
	p := &BookingManager{
		lg: lg.With().Caller().Logger(),
	}

	return p
}

func (m *BookingManager) PrepareBook(order *models.Order) (BookingID, error) {
	return BookingID{}, nil
}

func (m *BookingManager) AcceptBook(booking_id BookingID) error {
	return nil
}
