package manager

import (
	"applicationDesign/internal/logic/hotel"
	"applicationDesign/internal/models"
)

type BaseHotelManager interface {
	AddHotel(hotel *hotel.Hotel)
	PrepareBook(order models.Order) error
}
