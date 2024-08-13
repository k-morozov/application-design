package models

type AddHotel struct {
	HotelID string   `json:"hotel_id"`
	RoomsID []string `json:"rooms"`
}

func (o AddHotel) Validate() bool {
	if o.HotelID == "" {
		return false
	}

	if len(o.RoomsID) == 0 {
		return false
	}
	return true
}
