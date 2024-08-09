package logic

type Hotel struct {
}

var _ GuestHouseManager = &Hotel{}

func newHotel() GuestHouseManager {
	return &Hotel{}
}
