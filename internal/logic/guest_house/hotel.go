package guest_house

const (
	RoomFree RoomStatus = iota
	RoomBusy
	RoomLock
)

type HotelID string

type RoomID string

type RoomStatus int

type Room struct {
	RoomID RoomID
	Status RoomStatus
}

type Hotel struct {
	HotelID HotelID
	Rooms   map[RoomID]*Room
}

func (h *Hotel) AddRoom(roomID RoomID) {
	h.Rooms[roomID] = &Room{}
}

func NewHotel(hotelID HotelID) Hotel {
	return Hotel{
		HotelID: hotelID,
	}
}
