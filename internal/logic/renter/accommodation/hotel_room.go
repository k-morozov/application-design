package accommodation

import (
	"applicationDesign/internal/utils"
	"sort"
)

type HotelRoom struct {
	RoomID                TAccommodationID
	FreeRoomIntervals     []TIntervalAccommodation
	ReservedRoomIntervals []TIntervalAccommodation
}

var _ BaseAccommodation = &HotelRoom{}

func NewRoom(roomID TAccommodationID) HotelRoom {
	return HotelRoom{
		RoomID: roomID,
		FreeRoomIntervals: []TIntervalAccommodation{
			{
				From: utils.DefaultFromDateHotelAvailable,
				To:   utils.DefaultToDateHotelAvailable,
			},
		},
		ReservedRoomIntervals: []TIntervalAccommodation{},
	}
}

func (r *HotelRoom) ReserveByInterval(candidateInterval TIntervalAccommodation) bool {
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
				//r.FreeRoomIntervals = append(r.FreeRoomIntervals[:index], r.FreeRoomIntervals[index+1:]...)
				tmp := r.FreeRoomIntervals[len(r.FreeRoomIntervals)-1]
				r.FreeRoomIntervals[index] = tmp
				r.FreeRoomIntervals = r.FreeRoomIntervals[:len(r.FreeRoomIntervals)-1]
			}

			if splitInterval {
				r.FreeRoomIntervals[index].To = candidateInterval.From

				r.FreeRoomIntervals = append(r.FreeRoomIntervals, TIntervalAccommodation{
					From: candidateInterval.To,
					To:   oldTo,
				})
			}

			r.ReservedRoomIntervals = append(r.ReservedRoomIntervals, TIntervalAccommodation{
				From: candidateInterval.From,
				To:   candidateInterval.To,
			})

			r.sortIntervals()
			return true
		}
	}
	return false
}

func (r *HotelRoom) sortIntervals() {
	sort.Slice(r.FreeRoomIntervals, func(i, j int) bool {
		return r.FreeRoomIntervals[i].From.Before(r.FreeRoomIntervals[j].From)
	})
	sort.Slice(r.ReservedRoomIntervals, func(i, j int) bool {
		return r.ReservedRoomIntervals[i].From.Before(r.ReservedRoomIntervals[j].From)
	})
}
