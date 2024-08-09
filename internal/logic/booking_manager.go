package logic

import (
	"applicationDesign/internal/models"

	"github.com/rs/zerolog"
)

type BookingManager struct {
	lg                  zerolog.Logger
	book_queue          BookQueue
	guest_house_manager GuestHouseManager
}

var _ Manager = &BookingManager{}

func NewBookingManager(lg zerolog.Logger) Manager {
	p := &BookingManager{
		lg:                  lg.With().Caller().Logger(),
		book_queue:          newMemoryBookQueue(lg),
		guest_house_manager: newHotel(),
	}

	return p
}

func (m *BookingManager) PrepareBook(order *models.Order) (BookingID, error) {
	m.lg.Info().Msg("BookingManager: call PrepareBook")

	out := m.book_queue.Add(order)

	m.lg.Info().Msg("BookingManager: wait result")
	result := <-out
	if result.err != nil {
		m.lg.Error().Err(result.err).Msg("Failed prepare book")
		return BookingID{}, result.err
	}

	m.lg.Info().Any("book_id", result.id).Msg("BookingManager: PrepareBook finished")
	return result.id, nil
}

func (m *BookingManager) AcceptBook(booking_id BookingID) error {
	return nil
}
