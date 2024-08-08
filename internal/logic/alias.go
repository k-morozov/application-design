package logic

import (
	"github.com/google/uuid"
)

type BookingID uuid.UUID

func (id BookingID) String() string {
	return uuid.UUID(id).String()
}
