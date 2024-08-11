package logic

import "applicationDesign/internal/logic/renter"

type BaseBookingQueue interface {
	Add(order renter.HotelOrder) error
	Stop() error
	Worker()
}
