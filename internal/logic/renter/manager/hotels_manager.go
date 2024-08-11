package manager

import (
	"applicationDesign/internal/logic/renter"
	"applicationDesign/internal/logic/renter/accommodation"
	"applicationDesign/internal/models"
	"errors"
	"github.com/rs/zerolog"
	"sync"
)

type HotelManager struct {
	hotelsTableMutex sync.RWMutex
	hotelsTable      map[renter.TRenterID]*renter.Hotel
	lg               zerolog.Logger
}

var _ BaseRentersManager = &HotelManager{}

func NewGuestHouseManager(lg zerolog.Logger) BaseRentersManager {
	return &HotelManager{
		hotelsTable: make(map[renter.TRenterID]*renter.Hotel),
		lg:          lg,
	}
}

func (h *HotelManager) AddHotel(hotel *renter.Hotel) {
	h.hotelsTableMutex.Lock()
	defer h.hotelsTableMutex.Unlock()

	h.hotelsTable[hotel.RenterID] = hotel
}

func (h *HotelManager) PrepareBook(order models.Order) error {
	h.hotelsTableMutex.Lock()
	defer h.hotelsTableMutex.Unlock()

	hotel, ok := h.hotelsTable[renter.TRenterID(order.HotelID)]
	if !ok {
		return errors.New("renter not found for this order")
	}

	h.lg.Info().Str("room_id", order.RoomID).Msg("start reserve accommodation in renter")
	if err := hotel.ReserveAccommodation(accommodation.TAccommodationID(order.RoomID), accommodation.TIntervalAccommodation{From: order.From, To: order.To}); err != nil {
		return err
	}
	h.lg.Info().Str("room_id", order.RoomID).Msg("accommodation is reserved in renter")

	return nil
}
