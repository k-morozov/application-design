package guest_house

type HotelManager struct {
}

var _ GuestHouseManager = &HotelManager{}

func NewGuestHouseManager() GuestHouseManager {
	return &HotelManager{}
}

func (h *HotelManager) AddGuestHouse() {}

func (h *HotelManager) PrepareBook(order *HotelOrder) error {
	return nil
}
