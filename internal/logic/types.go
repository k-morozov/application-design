package logic

import (
	"github.com/google/uuid"
)

type TBookingID uuid.UUID

func (id TBookingID) String() string {
	return uuid.UUID(id).String()
}

func NewBookingID() TBookingID {
	return TBookingID(uuid.New())
}
