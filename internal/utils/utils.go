package utils

import "time"

var (
	DefaultFromDateHotelAvailable = Date(2030, 1, 1)
	DefaultToDateHotelAvailable   = Date(2030, 12, 31)
)

func ToDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
