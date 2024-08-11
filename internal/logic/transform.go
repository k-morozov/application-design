package logic

import (
	"applicationDesign/internal/logic/hotel"
	"applicationDesign/internal/models"
)

func transform(order models.Order) hotel.HotelOrder {
	return hotel.HotelOrder{
		ResultCh: make(chan error, 1),
		Order:    order,
	}
}
