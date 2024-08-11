package guest_house

import (
	"applicationDesign/internal/utils"
	"sort"
	"time"
)

type RoomID string

type RoomStatus int

type RoomInterval struct {
	From time.Time
	To   time.Time
}

type Room struct {
	RoomID                RoomID
	FreeRoomIntervals     []RoomInterval
	ReservedRoomIntervals []RoomInterval
}

func NewRoom(roomID RoomID) Room {
	return Room{
		RoomID: roomID,
		FreeRoomIntervals: []RoomInterval{
			{
				From: utils.Date(2030, 1, 1),
				To:   utils.Date(2030, 12, 31),
			},
		},
		ReservedRoomIntervals: []RoomInterval{},
	}
}

func (r *Room) ReserveByInterval(candidateInterval RoomInterval) bool {
	for index, interval := range r.FreeRoomIntervals {
		// early break

		resultFrom := interval.From.Compare(candidateInterval.From)
		resultTo := interval.To.Compare(candidateInterval.To)

		if resultFrom != 1 && resultTo != -1 {
			oldTo := r.FreeRoomIntervals[index].To

			splitInterval := true
			removeInterval := 0

			if r.FreeRoomIntervals[index].From == candidateInterval.From {
				r.FreeRoomIntervals[index].From = candidateInterval.To
				splitInterval = false
				removeInterval++
			}
			if r.FreeRoomIntervals[index].To == candidateInterval.To {
				r.FreeRoomIntervals[index].To = candidateInterval.From
				splitInterval = false
				removeInterval++
			}

			if removeInterval == 2 {
				r.FreeRoomIntervals = append(r.FreeRoomIntervals[:index], r.FreeRoomIntervals[index+1:]...)
			}

			if splitInterval {
				r.FreeRoomIntervals[index].To = candidateInterval.From

				r.FreeRoomIntervals = append(r.FreeRoomIntervals, RoomInterval{
					From: candidateInterval.To,
					To:   oldTo,
				})
			}

			r.ReservedRoomIntervals = append(r.ReservedRoomIntervals, RoomInterval{
				From: candidateInterval.From,
				To:   candidateInterval.To,
			})

			r.sortIntervals()
			return true
		}
	}
	return false
}

func (r *Room) sortIntervals() {
	sort.Slice(r.FreeRoomIntervals, func(i, j int) bool {
		return r.FreeRoomIntervals[i].From.Before(r.FreeRoomIntervals[j].From)
	})
	sort.Slice(r.ReservedRoomIntervals, func(i, j int) bool {
		return r.ReservedRoomIntervals[i].From.Before(r.ReservedRoomIntervals[j].From)
	})
}
