package renter

import "applicationDesign/internal/logic/renter/accommodation"

type TRenterID string

func (id TRenterID) String() string {
	return string(id)
}

type TBaseRenter interface {
	AddAccommodation(accommodationID accommodation.TAccommodationID)
	ReserveAccommodation(accommodationID accommodation.TAccommodationID, interval accommodation.TIntervalAccommodation) error
}
