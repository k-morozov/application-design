package guest_house

type GuestHouseManager interface {
	AddGuestHouse()
	PrepareBook(order *HotelOrder) error
}
