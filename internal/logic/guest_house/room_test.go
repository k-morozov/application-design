package guest_house

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoomIntervals(t *testing.T) {
	//lg := zerolog.Nop()
	//lg := log.NewLogger("debug")

	tests := []struct {
		name      string
		room      Room
		intervals []RoomInterval
		expected  []RoomInterval
	}{
		{
			name: "reserve one day",
			room: NewRoom("room1"),
			intervals: []RoomInterval{
				{
					From: Date(2030, 1, 2),
					To:   Date(2030, 1, 3),
				},
			},
			expected: []RoomInterval{
				{
					From: Date(2030, 1, 1),
					To:   Date(2030, 1, 2),
				},
				{
					From: Date(2030, 1, 3),
					To:   Date(2030, 12, 31),
				},
			},
		},
		{
			name: "reserve 2 times",
			room: NewRoom("room1"),
			intervals: []RoomInterval{
				{
					From: Date(2030, 1, 2),
					To:   Date(2030, 1, 15),
				},
				{
					From: Date(2030, 1, 16),
					To:   Date(2030, 2, 2),
				},
			},
			expected: []RoomInterval{
				{
					From: Date(2030, 1, 1),
					To:   Date(2030, 1, 2),
				},
				{
					From: Date(2030, 1, 15),
					To:   Date(2030, 1, 16),
				},
				{
					From: Date(2030, 2, 2),
					To:   Date(2030, 12, 31),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for _, interval := range test.intervals {
				reserved := test.room.ReserveByInterval(interval)
				assert.True(t, reserved)
			}
			assert.Equal(t, test.expected, test.room.FreeRoomIntervals)
		})
	}
}
