package accommodation

import (
	"applicationDesign/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoomIntervals(t *testing.T) {
	tests := []struct {
		name                      string
		room                      BaseAccommodation
		intervals                 []TIntervalAccommodation
		expectedFreeIntervals     []TIntervalAccommodation
		expectedReservedIntervals []TIntervalAccommodation
		reserved                  bool
	}{
		{
			name: "reserve one day",
			room: NewRoom("room1"),
			intervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 2),
					To:   utils.Date(2030, 1, 3),
				},
			},
			expectedFreeIntervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 1),
					To:   utils.Date(2030, 1, 2),
				},
				{
					From: utils.Date(2030, 1, 3),
					To:   utils.Date(2030, 12, 31),
				},
			},
			expectedReservedIntervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 2),
					To:   utils.Date(2030, 1, 3),
				},
			},
			reserved: true,
		},
		{
			name: "reserve 2 times",
			room: NewRoom("room1"),
			intervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 2),
					To:   utils.Date(2030, 1, 15),
				},
				{
					From: utils.Date(2030, 1, 16),
					To:   utils.Date(2030, 2, 2),
				},
			},
			expectedFreeIntervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 1),
					To:   utils.Date(2030, 1, 2),
				},
				{
					From: utils.Date(2030, 1, 15),
					To:   utils.Date(2030, 1, 16),
				},
				{
					From: utils.Date(2030, 2, 2),
					To:   utils.Date(2030, 12, 31),
				},
			},
			expectedReservedIntervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 2),
					To:   utils.Date(2030, 1, 15),
				},
				{
					From: utils.Date(2030, 1, 16),
					To:   utils.Date(2030, 2, 2),
				},
			},
			reserved: true,
		},
		{
			name: "reserve from begin",
			room: NewRoom("room1"),
			intervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 1),
					To:   utils.Date(2030, 1, 3),
				},
			},
			expectedFreeIntervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 3),
					To:   utils.Date(2030, 12, 31),
				},
			},
			expectedReservedIntervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 1),
					To:   utils.Date(2030, 1, 3),
				},
			},
			reserved: true,
		},
		{
			name: "reserve from begin many",
			room: NewRoom("room1"),
			intervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 3),
					To:   utils.Date(2030, 1, 10),
				},
				{
					From: utils.Date(2030, 1, 1),
					To:   utils.Date(2030, 1, 3),
				},
			},
			expectedFreeIntervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 10),
					To:   utils.Date(2030, 12, 31),
				},
			},
			expectedReservedIntervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 1),
					To:   utils.Date(2030, 1, 3),
				},
				{
					From: utils.Date(2030, 1, 3),
					To:   utils.Date(2030, 1, 10),
				},
			},
			reserved: true,
		},
		{
			name: "reserve to end",
			room: NewRoom("room1"),
			intervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 12, 1),
					To:   utils.Date(2030, 12, 31),
				},
			},
			expectedFreeIntervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 1),
					To:   utils.Date(2030, 12, 1),
				},
			},
			expectedReservedIntervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 12, 1),
					To:   utils.Date(2030, 12, 31),
				},
			},
			reserved: true,
		},
		{
			name: "reserve out begin",
			room: NewRoom("room1"),
			intervals: []TIntervalAccommodation{
				{
					From: utils.Date(2029, 12, 31),
					To:   utils.Date(2030, 1, 1),
				},
			},
			expectedFreeIntervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 1),
					To:   utils.Date(2030, 12, 31),
				},
			},
			expectedReservedIntervals: []TIntervalAccommodation{},
			reserved:                  false,
		},
		{
			name: "reserve out end",
			room: NewRoom("room1"),
			intervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 12, 31),
					To:   utils.Date(2031, 1, 1),
				},
			},
			expectedFreeIntervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 1),
					To:   utils.Date(2030, 12, 31),
				},
			},
			expectedReservedIntervals: []TIntervalAccommodation{},
			reserved:                  false,
		},
		{
			name: "reserve all",
			room: NewRoom("room1"),
			intervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 1),
					To:   utils.Date(2030, 12, 31),
				},
			},
			expectedFreeIntervals: []TIntervalAccommodation{},
			expectedReservedIntervals: []TIntervalAccommodation{
				{
					From: utils.Date(2030, 1, 1),
					To:   utils.Date(2030, 12, 31),
				},
			},
			reserved: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, interval := range test.intervals {
				reserved := test.room.ReserveByInterval(interval)
				assert.Equal(t, test.reserved, reserved)
			}
			assert.Equal(t, test.expectedFreeIntervals, test.room.GetFreeIntervals())
			assert.Equal(t, test.expectedReservedIntervals, test.room.GetReservedIntervals())
		})
	}
}
