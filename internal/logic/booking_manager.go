package logic

import (
	"applicationDesign/internal/logic/guest_house"
	"applicationDesign/internal/models"
	"github.com/rs/zerolog"
	"sync"
)

type BookingManager struct {
	lg          zerolog.Logger
	bookQueue   BookQueue
	ordersMutex sync.Mutex
	orders      map[BookingID]*models.Order
}

var _ Manager = &BookingManager{}

func NewBookingManager(guestHouseManager guest_house.GuestHouseManager, workers int, lg zerolog.Logger) Manager {
	p := &BookingManager{
		lg:        lg.With().Caller().Logger(),
		bookQueue: newMemoryBookQueue(guestHouseManager, lg, workers),
		orders:    make(map[BookingID]*models.Order),
	}

	return p
}

func (m *BookingManager) PrepareBook(order *models.Order) (BookingID, error) {
	m.lg.Info().Msg("BookingManager: call PrepareBook")

	internalOrder := transform(order)
	orderDescriptorCh := internalOrder.ResultCh
	if err := m.bookQueue.Add(internalOrder); err != nil {
		m.lg.Error().Err(err).Msg("Failed add")
		return BookingID{}, err
	}

	m.lg.Info().Msg("BookingManager: wait resultCh")
	err := <-orderDescriptorCh
	if err != nil {
		m.lg.Error().Err(err).Msg("Failed prepare book")
		return BookingID{}, err
	}

	bookingID, err := m.SaveOrder(order)
	if err != nil {
		m.lg.Error().Err(err).Msg("Failed save order")
		return BookingID{}, err
	}

	m.lg.Info().Any("book_id", bookingID).Msg("BookingManager: PrepareBook finished")
	return bookingID, nil
}

func (m *BookingManager) AcceptBook(bookingID BookingID) error {
	return nil
}

func (m *BookingManager) SaveOrder(order *models.Order) (BookingID, error) {
	bookingID := NewBookingID()
	m.ordersMutex.Lock()
	defer m.ordersMutex.Unlock()

	m.orders[bookingID] = order

	m.lg.Info().Any("bookingID", bookingID).Msg("BookingManager: successfully generate bookingID and save order")

	return bookingID, nil
}
