package logic

import (
	"applicationDesign/internal/log"
	"applicationDesign/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMemoryBookingQueue(t *testing.T) {
	// lg := zerolog.Nop()
	lg := log.NewLogger("debug")

	tests := []struct {
		name   string
		orders []models.Order
	}{
		{
			name: "simple",
			orders: []models.Order{
				{
					HotelID:   "hotel1",
					RoomID:    "room1",
					UserEmail: "a@a",
					From:      time.Time{},
					To:        time.Time{},
				},
				{
					HotelID:   "hotel1",
					RoomID:    "room2",
					UserEmail: "a@a",
					From:      time.Time{},
					To:        time.Time{},
				},
				{
					HotelID:   "hotel1",
					RoomID:    "room3",
					UserEmail: "a@a",
					From:      time.Time{},
					To:        time.Time{},
				},
				{
					HotelID:   "hotel1",
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
			q := newMemoryBookQueue(lg, 2)

			var results []chan ResultPrepareBook

			for _, order := range test.orders {
				internalOrder := transform(&order)
				q.Add(internalOrder)
				results = append(results, internalOrder.resultCh)
			}

			for _, resultCh := range results {
				lg.Info().Msg("wait result")
				result := <-resultCh
				lg.Info().Msg("result ready")

				close(resultCh)

				assert.Equal(t, nil, result.err)
				assert.NotEqual(t, "", result.id.String())
			}

			q.Stop()

		})
	}
}
