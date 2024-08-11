package manager

import (
	"applicationDesign/internal/logic/hotel"
	"applicationDesign/internal/logic/hotel/accommodation"
	"applicationDesign/internal/models"
	"errors"
	"github.com/rs/zerolog"
	"sync"
)

type HotelManager struct {
	hotelsTableMutex sync.RWMutex
	hotelsTable      map[hotel.HotelID]*hotel.Hotel
	lg               zerolog.Logger
}

var _ BaseHotelManager = &HotelManager{}

func NewGuestHouseManager(lg zerolog.Logger) BaseHotelManager {
	return &HotelManager{
		hotelsTable: make(map[hotel.HotelID]*hotel.Hotel),
		lg:          lg,
	}
}

func (h *HotelManager) AddHotel(hotel *hotel.Hotel) {
	h.hotelsTableMutex.Lock()
	defer h.hotelsTableMutex.Unlock()

	h.hotelsTable[hotel.HotelID] = hotel
}

func (h *HotelManager) PrepareBook(order models.Order) error {
	h.hotelsTableMutex.Lock()
	defer h.hotelsTableMutex.Unlock()

	hotel, ok := h.hotelsTable[hotel.HotelID(order.HotelID)]
	if !ok {
		return errors.New("hotel not found for this order")
	}

	h.lg.Info().Str("room_id", order.RoomID).Msg("start reserve accommodation in hotel")
	if err := hotel.ReserveRoom(accommodation.AccommodationID(order.RoomID), accommodation.IntervalAccommodation{From: order.From, To: order.To}); err != nil {
		return err
	}
	h.lg.Info().Str("room_id", order.RoomID).Msg("accommodation is reserved in hotel")

	return nil
}
