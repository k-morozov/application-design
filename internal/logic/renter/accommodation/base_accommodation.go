package accommodation

import "time"

type TAccommodationID string

type TStatusAccommodation int

type TIntervalAccommodation struct {
	From time.Time
	To   time.Time
}

type BaseAccommodation interface {
	ReserveByInterval(candidateInterval TIntervalAccommodation) bool
}
