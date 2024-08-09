package logic

import (
	"applicationDesign/internal/models"
	"github.com/rs/zerolog"
)

type BookingManager struct {
	lg                zerolog.Logger
	bookQueue         BookQueue
	guestHouseManager GuestHouseManager
}

var _ Manager = &BookingManager{}

func NewBookingManager(lg zerolog.Logger, workers int) Manager {
	p := &BookingManager{
		lg:                lg.With().Caller().Logger(),
		bookQueue:         newMemoryBookQueue(lg, workers),
		guestHouseManager: newHotel(),
	}

	return p
}

func (m *BookingManager) PrepareBook(order *models.Order) (BookingID, error) {
	m.lg.Info().Msg("BookingManager: call PrepareBook")

	internalOrder := transform(order)
	m.bookQueue.Add(internalOrder)

	m.lg.Info().Msg("BookingManager: wait resultCh")
	result := <-internalOrder.resultCh
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
