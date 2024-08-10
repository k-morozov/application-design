package guest_house

import (
	"applicationDesign/internal/models"
	"errors"
	"github.com/rs/zerolog"
	"sync"
)

type HotelManager struct {
	hotelsTableMutex sync.RWMutex
	hotelsTable      map[HotelID]*Hotel
	lg               zerolog.Logger
}

var _ GuestHouseManager = &HotelManager{}

func NewGuestHouseManager(lg zerolog.Logger) GuestHouseManager {
	return &HotelManager{
		hotelsTable: make(map[HotelID]*Hotel),
		lg:          lg,
	}
}

func (h *HotelManager) AddGuestHouse(hotel *Hotel) {
	h.hotelsTableMutex.Lock()
	defer h.hotelsTableMutex.Unlock()

	h.hotelsTable[hotel.HotelID] = hotel
}

func (h *HotelManager) PrepareBook(order models.Order) error {
	h.hotelsTableMutex.Lock()
	defer h.hotelsTableMutex.Unlock()

	hotel, ok := h.hotelsTable[HotelID(order.HotelID)]
	if !ok {
		return errors.New("hotel not found for this order")
	}

	h.lg.Info().Str("room_id", order.RoomID).Msg("start reserve room in hotel")
	if err := hotel.ReserveRoom(RoomID(order.RoomID), RoomInterval{From: order.From, To: order.To}); err != nil {
		return err
	}
	h.lg.Info().Str("room_id", order.RoomID).Msg("room is reserved in hotel")

	return nil
}
