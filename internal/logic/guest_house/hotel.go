package guest_house

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"sync"
)

const (
	RoomFree RoomStatus = iota
	RoomReserve
	RoomBusy
)

type HotelID string

func (id HotelID) String() string {
	return string(id)
}

type Hotel struct {
	HotelID    HotelID
	Rooms      map[RoomID]*Room
	roomsMutex sync.RWMutex
	lg         zerolog.Logger
}

func (h *Hotel) AddRoom(roomID RoomID) {
	h.roomsMutex.Lock()
	defer h.roomsMutex.Unlock()

	r := NewRoom(roomID)
	h.Rooms[roomID] = &r
}

func NewHotel(hotelID HotelID, lg zerolog.Logger) Hotel {
	return Hotel{
		HotelID: hotelID,
		Rooms:   map[RoomID]*Room{},
		lg:      lg,
	}
}

func (h *Hotel) ReserveRoom(roomID RoomID, interval RoomInterval) error {
	h.roomsMutex.Lock()
	defer h.roomsMutex.Unlock()

	if h.Rooms[roomID] == nil {
		return errors.New("room is not exists")
	}

	h.lg.Info().Str("room_id", string(roomID)).Any("rooms", h.Rooms).Msg("Status all rooms in hotel")

	//if h.Rooms[roomID].Status != RoomFree {
	//	return fmt.Errorf("room with id=%v is not free: %v, should be %v", roomID, h.Rooms[roomID].Status, RoomFree)
	//}

	if !h.Rooms[roomID].ReserveByInterval(interval) {
		return fmt.Errorf("room with id=%v has already been reserved", roomID)
	}

	//h.Rooms[roomID].Status = RoomReserve

	return nil
}
