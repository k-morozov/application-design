package models

import "time"

type Order struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

func (o Order) Validate() bool {
	if o.HotelID == "" {
		return false
	}

	if o.RoomID == "" {
		return false
	}

	if o.UserEmail == "" {
		return false
	}

	var DefaultTime time.Time

	if o.From.String() == DefaultTime.String() {
		return false
	}

	if o.To.String() == DefaultTime.String() {
		return false
	}

	// validate email, time ...

	return true
}
