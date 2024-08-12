package rental

import "applicationDesign/internal/logic/rental/accommodation"

type TRentalID string

func (id TRentalID) String() string {
	return string(id)
}

type TBaseRental interface {
	GetRentalID() TRentalID
	GetTableAccommodation() map[accommodation.TAccommodationID]accommodation.BaseAccommodation
	AddAccommodation(accommodationID accommodation.TAccommodationID)
	ReserveAccommodation(accommodationID accommodation.TAccommodationID, interval accommodation.TIntervalAccommodation) error
}
