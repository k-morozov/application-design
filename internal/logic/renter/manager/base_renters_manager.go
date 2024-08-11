package manager

import (
	"applicationDesign/internal/logic/renter"
	"applicationDesign/internal/models"
)

type BaseRentersManager interface {
	AddHotel(hotel *renter.Hotel)
	PrepareBook(order models.Order) error
}
