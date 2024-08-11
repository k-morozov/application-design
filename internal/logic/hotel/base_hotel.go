package hotel

import "applicationDesign/internal/logic/hotel/accommodation"

type HotelID string

func (id HotelID) String() string {
	return string(id)
}

type BaseHotel interface {
	AddRoom(roomID accommodation.AccommodationID)
	ReserveRoom(roomID accommodation.AccommodationID, interval accommodation.IntervalAccommodation) error
}
