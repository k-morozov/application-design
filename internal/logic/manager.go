package logic

import "applicationDesign/internal/models"

type Manager interface {
	PrepareBook(order models.Order) (BookingID, error)
	AcceptBook(bookingId BookingID) error
}
