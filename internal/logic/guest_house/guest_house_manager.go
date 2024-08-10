package guest_house

import "applicationDesign/internal/models"

type GuestHouseManager interface {
	AddGuestHouse(hotel *Hotel)
	PrepareBook(order models.Order) error
}
