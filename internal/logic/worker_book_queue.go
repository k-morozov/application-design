package logic

import (
	"applicationDesign/internal/logic/guest_house"
	"github.com/rs/zerolog"
	"sync"
)

type WorkerBookQueue struct {
	guestHouseManager guest_house.GuestHouseManager
	lg                zerolog.Logger
	ordersQueue       chan *guest_house.HotelOrder
	wg                sync.WaitGroup
}

var _ BookQueue = &WorkerBookQueue{}

func newMemoryBookQueue(guestHouseManager guest_house.GuestHouseManager, lg zerolog.Logger, workers int) BookQueue {
	result := &WorkerBookQueue{
		guestHouseManager: guestHouseManager,
		lg:                lg,
		ordersQueue:       make(chan *guest_house.HotelOrder),
	}
	for w := 0; w < workers; w++ {
		result.lg.Debug().Msg("Add worker queue")
		result.wg.Add(1)
		go result.Worker()
	}
	return result
}

func (q *WorkerBookQueue) Add(order *guest_house.HotelOrder) error {
	q.lg.Info().Msg("Add: start")

	// @todo if channel is close?
	q.ordersQueue <- order

	q.lg.Info().Msg("Add: finish")

	return nil
}

func (q *WorkerBookQueue) Stop() error {
	q.lg.Debug().Msg("Stop queue")
	close(q.ordersQueue)

	q.lg.Debug().Msg("Wait close worker")
	q.wg.Wait()

	return nil
}

func (q *WorkerBookQueue) Worker() {
	for order := range q.ordersQueue {
		q.lg.Debug().Msg("worker has an order.")

		result := q.guestHouseManager.PrepareBook(order)
		order.ResultCh <- result

		q.lg.Debug().Msg("worker send result in channel.")
	}
	q.lg.Debug().Msg("worker has done.")
	q.wg.Done()
}
