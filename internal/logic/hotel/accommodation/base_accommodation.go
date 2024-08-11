package accommodation

import "time"

type AccommodationID string

type StatusAccommodation int

type IntervalAccommodation struct {
	From time.Time
	To   time.Time
}

type BaseAccommodation interface {
	ReserveByInterval(candidateInterval IntervalAccommodation) bool
}
