package rental

import (
	"applicationDesign/internal/logic/rental/accommodation"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"sync"
)

type Hotel struct {
	RenterID   TRentalID
	Rooms      map[accommodation.TAccommodationID]*accommodation.HotelRoom
	roomsMutex sync.RWMutex
	lg         zerolog.Logger
}

var _ TBaseRental = &Hotel{}

func NewHotel(renterID TRentalID, lg zerolog.Logger) Hotel {
	return Hotel{
		RenterID: renterID,
		Rooms:    map[accommodation.TAccommodationID]*accommodation.HotelRoom{},
		lg:       lg,
	}
}

func (h *Hotel) GetRentalID() TRentalID {
	return h.RenterID
}

func (h *Hotel) AddAccommodation(roomID accommodation.TAccommodationID) {
	h.roomsMutex.Lock()
	defer h.roomsMutex.Unlock()

	r := accommodation.NewRoom(roomID)
	h.Rooms[roomID] = &r
}

func (h *Hotel) ReserveAccommodation(roomID accommodation.TAccommodationID, interval accommodation.TIntervalAccommodation) error {
	h.roomsMutex.Lock()
	defer h.roomsMutex.Unlock()

	if h.Rooms[roomID] == nil {
		return errors.New("accommodation is not exists")
	}

	h.lg.Info().Str("room_id", string(roomID)).Any("rooms", h.Rooms).Msg("Status all rooms in rental")

	if !h.Rooms[roomID].ReserveByInterval(interval) {
		return fmt.Errorf("accommodation with id=%v has already been reserved", roomID)
	}

	return nil
}
