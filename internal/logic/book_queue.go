package logic

import "applicationDesign/internal/models"

type ResultPrepareBook struct {
	err error
	id  BookingID
}

type BookQueue interface {
	Add(order *models.Order) <-chan ResultPrepareBook
}
