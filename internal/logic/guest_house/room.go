package guest_house

import "time"

type RoomID string

type RoomStatus int

type RoomInterval struct {
	From time.Time
	To   time.Time
}

type Room struct {
	RoomID            RoomID
	Status            RoomStatus
	FreeRoomIntervals []RoomInterval
}

func NewRoom(roomID RoomID) Room {
	return Room{
		RoomID: roomID,
		FreeRoomIntervals: []RoomInterval{
			{
				From: Date(2030, 1, 1),
				To:   Date(2030, 12, 31),
			},
		},
	}
}

func (r *Room) ReserveByInterval(candidateInterval RoomInterval) bool {
	for index, interval := range r.FreeRoomIntervals {
		resultFrom := interval.From.Compare(candidateInterval.From)
		resultTo := interval.To.Compare(candidateInterval.To)

		if resultFrom != 1 && resultTo != -1 {
			oldTo := r.FreeRoomIntervals[index].To
			r.FreeRoomIntervals[index].To = candidateInterval.From

			r.FreeRoomIntervals = append(r.FreeRoomIntervals, RoomInterval{
				From: candidateInterval.To,
				To:   oldTo,
			})
			return true
		}
	}
	return false
}

func ToDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
