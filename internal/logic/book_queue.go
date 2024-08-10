package logic

import "applicationDesign/internal/logic/guest_house"

type BookQueue interface {
	Add(order guest_house.HotelOrder) error
	Stop() error
	Worker()
}
