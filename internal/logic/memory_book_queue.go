package logic

import (
	"applicationDesign/internal/models"

	"github.com/rs/zerolog"
)

type MemoryBookQueue struct {
	lg           zerolog.Logger
	orders_queue chan InternalOrder
	// stop workers
}

var _ BookQueue = &MemoryBookQueue{}

func newMemoryBookQueue(lg zerolog.Logger) BookQueue {
	result := &MemoryBookQueue{
		lg:           lg,
		orders_queue: make(chan InternalOrder),
	}
	for w := 0; w < 2; w++ {
		result.lg.Debug().Msg("Add worker queue")
		go worker(result.orders_queue)
	}
	return result
}

func (q *MemoryBookQueue) Add(order *models.Order) <-chan ResultPrepareBook {
	q.lg.Debug().Msg("Add order")

	result_channel := make(chan ResultPrepareBook)
	internal_order := InternalOrder{
		result: result_channel,
		order:  order,
	}

	q.orders_queue <- internal_order

	return result_channel
}

func worker(orders <-chan InternalOrder) {
	for order := range orders {
		result := ResultPrepareBook{
			err: nil,
			id:  NewBookingID(),
		}
		order.result <- result
	}
}
