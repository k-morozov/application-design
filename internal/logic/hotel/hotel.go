package hotel

import (
	"applicationDesign/internal/logic/hotel/accommodation"
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"sync"
)

type Hotel struct {
	HotelID    HotelID
	Rooms      map[accommodation.AccommodationID]*accommodation.HotelRoom
	roomsMutex sync.RWMutex
	lg         zerolog.Logger
}

var _ BaseHotel = &Hotel{}

func NewHotel(hotelID HotelID, lg zerolog.Logger) Hotel {
	return Hotel{
		HotelID: hotelID,
		Rooms:   map[accommodation.AccommodationID]*accommodation.HotelRoom{},
		lg:      lg,
	}
}

func (h *Hotel) AddRoom(roomID accommodation.AccommodationID) {
	h.roomsMutex.Lock()
	defer h.roomsMutex.Unlock()

	r := accommodation.NewRoom(roomID)
	h.Rooms[roomID] = &r
}

func (h *Hotel) ReserveRoom(roomID accommodation.AccommodationID, interval accommodation.IntervalAccommodation) error {
	h.roomsMutex.Lock()
	defer h.roomsMutex.Unlock()

	if h.Rooms[roomID] == nil {
		return errors.New("accommodation is not exists")
	}

	h.lg.Info().Str("room_id", string(roomID)).Any("rooms", h.Rooms).Msg("Status all rooms in hotel")

	if !h.Rooms[roomID].ReserveByInterval(interval) {
		return fmt.Errorf("accommodation with id=%v has already been reserved", roomID)
	}

	return nil
}
