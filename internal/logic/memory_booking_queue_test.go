package logic

import (
	"applicationDesign/internal/log"
	"applicationDesign/internal/logic/guest_house"
	"applicationDesign/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMemoryBookingQueue(t *testing.T) {
	// lg := zerolog.Nop()
	lg := log.NewLogger("debug")

	HotelID := guest_house.HotelID("hotel1")
	tests := []struct {
		name   string
		orders []models.Order
	}{
		{
			name: "simple",
			orders: []models.Order{
				{
					HotelID:   HotelID.String(),
					RoomID:    "room1",
					UserEmail: "a@a",
					From:      time.Time{},
					To:        time.Time{},
				},
				{
					HotelID:   HotelID.String(),
					RoomID:    "room2",
					UserEmail: "b@b",
					From:      time.Time{},
					To:        time.Time{},
				},
				{
					HotelID:   HotelID.String(),
					RoomID:    "room3",
					UserEmail: "c@c",
					From:      time.Time{},
					To:        time.Time{},
				},
				{
					HotelID:   HotelID.String(),
					RoomID:    "room4",
					UserEmail: "a@a",
					From:      time.Time{},
					To:        time.Time{},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := guest_house.NewGuestHouseManager(lg)

			hotel := guest_house.NewHotel(HotelID, lg)
			for _, order := range test.orders {
				hotel.AddRoom(guest_house.RoomID(order.RoomID))
			}

			g.AddGuestHouse(&hotel)

			q := newMemoryBookQueue(g, lg, 2)

			var results []chan error

			for _, order := range test.orders {
				internalOrder := transform(order)
				_ = q.Add(internalOrder)
				results = append(results, internalOrder.ResultCh)
			}

			for _, resultCh := range results {
				err := <-resultCh

				close(resultCh)

				assert.Equal(t, nil, err)
				//assert.NotEqual(t, "", result.bookingID.String())
			}

			_ = q.Stop()

		})
	}
}
