package logic

import (
	"applicationDesign/internal/logic/guest_house"
	"applicationDesign/internal/models"
)

func transform(order models.Order) guest_house.HotelOrder {
	return guest_house.HotelOrder{
		ResultCh: make(chan error, 1),
		Order:    order,
	}
}
