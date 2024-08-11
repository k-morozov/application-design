package logic

import (
	"applicationDesign/internal/logic/renter/manager"
	"applicationDesign/internal/models"
	"github.com/rs/zerolog"
	"sync"
)

type BookingManager struct {
	lg          zerolog.Logger
	bookQueue   BaseBookingQueue
	ordersMutex sync.Mutex
	orders      map[TBookingID]models.Order
}

var _ BaseBookingManager = &BookingManager{}

func NewBookingManager(guestHouseManager manager.BaseRentersManager, workers int, lg zerolog.Logger) BaseBookingManager {
	p := &BookingManager{
		lg:        lg.With().Caller().Logger(),
		bookQueue: NewInMemoryBookingQueue(guestHouseManager, lg, workers),
		orders:    make(map[TBookingID]models.Order),
	}

	return p
}

func (m *BookingManager) PrepareBook(order models.Order) (TBookingID, error) {
	m.lg.Info().Msg("BookingManager: call PrepareBook")

	internalOrder := transform(order)
	orderDescriptorCh := internalOrder.ResultCh
	if err := m.bookQueue.Add(internalOrder); err != nil {
		m.lg.Error().Err(err).Msg("Failed add")
		return TBookingID{}, err
	}

	m.lg.Info().Msg("BookingManager: wait resultCh")
	err := <-orderDescriptorCh
	if err != nil {
		m.lg.Error().Err(err).Msg("Failed prepare book")
		return TBookingID{}, err
	}

	bookingID, err := m.SaveOrder(order)
	if err != nil {
		m.lg.Error().Err(err).Msg("Failed save order")
		return TBookingID{}, err
	}

	m.lg.Info().Any("book_id", bookingID).Msg("BookingManager: PrepareBook finished")
	return bookingID, nil
}

func (m *BookingManager) AcceptBook(bookingID TBookingID) error {
	return nil
}

func (m *BookingManager) SaveOrder(order models.Order) (TBookingID, error) {
	bookingID := NewBookingID()
	m.ordersMutex.Lock()
	defer m.ordersMutex.Unlock()

	m.orders[bookingID] = order

	m.lg.Info().Any("bookingID", bookingID).Msg("BookingManager: successfully generate bookingID and save order")

	return bookingID, nil
}
