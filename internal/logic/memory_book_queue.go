package logic

import (
	"github.com/rs/zerolog"
	"sync"
)

type MemoryBookQueue struct {
	lg          zerolog.Logger
	ordersQueue chan *InternalOrder
	wg          sync.WaitGroup
}

var _ BookQueue = &MemoryBookQueue{}

func newMemoryBookQueue(lg zerolog.Logger, workers int) BookQueue {
	result := &MemoryBookQueue{
		lg:          lg,
		ordersQueue: make(chan *InternalOrder),
	}
	for w := 0; w < workers; w++ {
		result.lg.Debug().Msg("Add worker queue")
		result.wg.Add(1)
		go result.worker()
	}
	return result
}

func (q *MemoryBookQueue) Add(order *InternalOrder) {
	q.lg.Info().Msg("Add: start")

	q.ordersQueue <- order

	q.lg.Info().Msg("Add: finish")
}

func (q *MemoryBookQueue) Stop() {
	q.lg.Debug().Msg("Stop queue")
	close(q.ordersQueue)

	q.lg.Debug().Msg("Wait close worker")
	q.wg.Wait()
}

func (q *MemoryBookQueue) worker() {
	for order := range q.ordersQueue {
		q.lg.Debug().Msg("worker has an order.")
		result := ResultPrepareBook{
			err: nil,
			id:  NewBookingID(),
		}
		order.resultCh <- result
		q.lg.Debug().Msg("worker send result in channel.")
	}
	q.lg.Debug().Msg("worker has done.")
	q.wg.Done()
}
