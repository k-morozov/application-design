package logic

import "applicationDesign/internal/logic/hotel"

type BaseBookingQueue interface {
	Add(order hotel.HotelOrder) error
	Stop() error
	Worker()
}
