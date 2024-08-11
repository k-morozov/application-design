package logic

import (
	"applicationDesign/internal/logic/renter"
	"applicationDesign/internal/logic/renter/manager"
	"github.com/rs/zerolog"
	"sync"
)

type InMemoryBookingQueue struct {
	guestHouseManager manager.BaseRentersManager
	lg                zerolog.Logger
	ordersQueue       chan renter.HotelOrder
	wg                sync.WaitGroup
}

var _ BaseBookingQueue = &InMemoryBookingQueue{}

func NewInMemoryBookingQueue(guestHouseManager manager.BaseRentersManager, lg zerolog.Logger, workers int) BaseBookingQueue {
	result := &InMemoryBookingQueue{
		guestHouseManager: guestHouseManager,
		lg:                lg,
		ordersQueue:       make(chan renter.HotelOrder),
	}
	for w := 0; w < workers; w++ {
		result.lg.Debug().Msg("Add worker queue")
		result.wg.Add(1)
		go result.Worker()
	}
	return result
}

func (q *InMemoryBookingQueue) Add(order renter.HotelOrder) error {
	q.lg.Info().Any("order", order.Order).Msg("add order to booking queue")

	// @todo if channel is close?
	q.ordersQueue <- order
	return nil
}

func (q *InMemoryBookingQueue) Stop() error {
	q.lg.Debug().Msg("Stop queue")
	close(q.ordersQueue)

	q.lg.Debug().Msg("Wait close worker")
	q.wg.Wait()

	return nil
}

func (q *InMemoryBookingQueue) Worker() {
	for order := range q.ordersQueue {
		//q.lg.Debug().Any("order", order.Order).Msg("worker gets order from booking queue.")

		result := q.guestHouseManager.PrepareBook(order.Order)
		order.ResultCh <- result

		//q.lg.Debug().Msg("worker send result in channel.")
	}
	q.lg.Debug().Msg("worker has done.")
	q.wg.Done()
}
