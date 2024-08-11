package logic

import (
	"applicationDesign/internal/logic/rental"
	"applicationDesign/internal/models"
)

func transform(order models.Order) rental.HotelOrder {
	return rental.HotelOrder{
		ResultCh: make(chan error, 1),
		Order:    order,
	}
}
