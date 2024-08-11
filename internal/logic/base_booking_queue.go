package logic

import "applicationDesign/internal/logic/rental"

type BaseBookingQueue interface {
	Add(order rental.HotelOrder) error
	Stop() error
	Worker()
}
