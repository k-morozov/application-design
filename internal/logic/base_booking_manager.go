package logic

import "applicationDesign/internal/models"

type BaseBookingManager interface {
	PrepareBook(order models.Order) (BookingID, error)
	AcceptBook(bookingId BookingID) error
}
