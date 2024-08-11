package logic

import (
	"applicationDesign/internal/logic/rental"
	"applicationDesign/internal/logic/rental/accommodation"
	"applicationDesign/internal/logic/rental/rental_manager"
	"applicationDesign/internal/models"
	"applicationDesign/internal/utils"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleBook(t *testing.T) {
	lg := zerolog.Nop()
	//lg := log.NewLogger("debug")

	HotelID := rental.TRentalID("hotel1")
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
					From:      utils.Date(2030, 1, 11),
					To:        utils.Date(2030, 1, 21),
				},
				{
					HotelID:   HotelID.String(),
					RoomID:    "room2",
					UserEmail: "b@b",
					From:      utils.Date(2030, 1, 11),
					To:        utils.Date(2030, 1, 21),
				},
				{
					HotelID:   HotelID.String(),
					RoomID:    "room3",
					UserEmail: "c@c",
					From:      utils.Date(2030, 1, 11),
					To:        utils.Date(2030, 1, 21),
				},
				{
					HotelID:   HotelID.String(),
					RoomID:    "room4",
					UserEmail: "a@a",
					From:      utils.Date(2030, 1, 11),
					To:        utils.Date(2030, 1, 21),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := rental_manager.NewGuestHouseManager(lg)

			testHotel := rental.NewHotel(HotelID, lg)
			for _, order := range test.orders {
				testHotel.AddAccommodation(accommodation.TAccommodationID(order.RoomID))
			}

			g.AddLandlord(&testHotel)

			q := NewInMemoryBookingQueue(g, lg, 2)

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
			}

			_ = q.Stop()

			assert.Equal(t, len(test.orders), len(testHotel.Rooms))
			for _, order := range test.orders {
				room, ok := testHotel.Rooms[accommodation.TAccommodationID(order.RoomID)]

				assert.True(t, ok)

				assert.Equal(t, room.FreeRoomIntervals, []accommodation.TIntervalAccommodation{
					{
						From: utils.Date(2030, 1, 1),
						To:   utils.Date(2030, 1, 11),
					},
					{
						From: utils.Date(2030, 1, 21),
						To:   utils.Date(2030, 12, 31),
					},
				})
			}
		})
	}
}

func TestBookOneRoom(t *testing.T) {
	lg := zerolog.Nop()
	//lg := log.NewLogger("debug")

	HotelID := rental.TRentalID("hotel1")
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
					From:      utils.Date(2030, 4, 11),
					To:        utils.Date(2030, 4, 21),
				},
				{
					HotelID:   HotelID.String(),
					RoomID:    "room1",
					UserEmail: "b@b",
					From:      utils.Date(2030, 4, 11),
					To:        utils.Date(2030, 4, 21),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g := rental_manager.NewGuestHouseManager(lg)

			testHotel := rental.NewHotel(HotelID, lg)
			for _, order := range test.orders {
				testHotel.AddAccommodation(accommodation.TAccommodationID(order.RoomID))
			}

			g.AddLandlord(&testHotel)

			q := NewInMemoryBookingQueue(g, lg, 2)

			var results []chan error

			for _, order := range test.orders {
				internalOrder := transform(order)
				_ = q.Add(internalOrder)
				results = append(results, internalOrder.ResultCh)
			}

			var flag bool
			for _, resultCh := range results {
				err := <-resultCh

				close(resultCh)

				if flag == false && err == nil {
					flag = true
					continue
				}

				assert.EqualError(t, err, "accommodation with id=room1 has already been reserved")
			}

			assert.True(t, flag)

			_ = q.Stop()

			assert.Equal(t, 1, len(testHotel.Rooms))
			for _, room := range testHotel.Rooms {
				assert.Equal(t, room.FreeRoomIntervals, []accommodation.TIntervalAccommodation{
					{
						From: utils.Date(2030, 1, 1),
						To:   utils.Date(2030, 4, 11),
					},
					{
						From: utils.Date(2030, 4, 21),
						To:   utils.Date(2030, 12, 31),
					},
				})
			}
		})
	}
}
