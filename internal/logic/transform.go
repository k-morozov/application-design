package logic

import "applicationDesign/internal/models"

func transform(order *models.Order) *InternalOrder {
	return &InternalOrder{
		resultCh: make(chan ResultPrepareBook, 1),
		order:    order,
	}
}
