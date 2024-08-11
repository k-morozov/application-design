package logic

import "applicationDesign/internal/models"

type BaseBookingManager interface {
	PrepareBook(order models.Order) (TBookingID, error)
	AcceptBook(bookingId TBookingID) error
}
