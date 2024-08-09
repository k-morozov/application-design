package logic

import "applicationDesign/internal/models"

type InternalOrder struct {
	result chan ResultPrepareBook
	order  *models.Order
}
