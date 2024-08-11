package logic

import (
	"applicationDesign/internal/logic/renter"
	"applicationDesign/internal/models"
)

func transform(order models.Order) renter.HotelOrder {
	return renter.HotelOrder{
		ResultCh: make(chan error, 1),
		Order:    order,
	}
}
