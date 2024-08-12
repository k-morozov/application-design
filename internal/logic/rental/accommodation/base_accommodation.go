package accommodation

import "time"

type TAccommodationID string

type TStatusAccommodation int

type TIntervalAccommodation struct {
	From time.Time
	To   time.Time
}

type BaseAccommodation interface {
	GetFreeIntervals() []TIntervalAccommodation
	GetReservedIntervals() []TIntervalAccommodation
	ReserveByInterval(candidateInterval TIntervalAccommodation) bool
}
