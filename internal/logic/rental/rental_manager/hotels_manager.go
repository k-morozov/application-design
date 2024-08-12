package rental_manager

import (
	"applicationDesign/internal/logic/rental"
	"applicationDesign/internal/logic/rental/accommodation"
	"applicationDesign/internal/models"
	"errors"
	"github.com/rs/zerolog"
	"sync"
)

type HotelManager struct {
	hotelsTableMutex sync.RWMutex
	hotelsTable      map[rental.TRentalID]rental.TBaseRental
	lg               zerolog.Logger
}

var _ BaseRentalManager = &HotelManager{}

func NewHotelManager(lg zerolog.Logger) BaseRentalManager {
	return &HotelManager{
		hotelsTable: make(map[rental.TRentalID]rental.TBaseRental),
		lg:          lg,
	}
}

func (h *HotelManager) AddRental(hotel rental.TBaseRental) {
	h.hotelsTableMutex.Lock()
	defer h.hotelsTableMutex.Unlock()

	h.hotelsTable[hotel.GetRentalID()] = hotel
}

func (h *HotelManager) PrepareBook(order models.Order) error {
	h.hotelsTableMutex.Lock()
	defer h.hotelsTableMutex.Unlock()

	hotel, ok := h.hotelsTable[rental.TRentalID(order.HotelID)]
	if !ok {
		return errors.New("rental not found for this order")
	}

	h.lg.Info().Str("room_id", order.RoomID).Msg("start reserve accommodation in rental")
	if err := hotel.ReserveAccommodation(accommodation.TAccommodationID(order.RoomID), accommodation.TIntervalAccommodation{From: order.From, To: order.To}); err != nil {
		return err
	}
	h.lg.Info().Str("room_id", order.RoomID).Msg("accommodation is reserved in rental")

	return nil
}
