package rental_manager

import (
	"applicationDesign/internal/logic/rental"
	"applicationDesign/internal/models"
)

type BaseRentalManager interface {
	AddRental(renter rental.TBaseRental)
	PrepareBook(order models.Order) error
}
