package logic

import "applicationDesign/internal/models"

type InternalOrder struct {
	resultCh chan ResultPrepareBook
	order    *models.Order
}
