package rental

import (
	"applicationDesign/internal/models"
)

type ResultPrepareBook struct {
	err error
}

type HotelOrder struct {
	ResultCh chan error
	Order    models.Order
}
