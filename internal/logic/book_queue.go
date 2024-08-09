package logic

type ResultPrepareBook struct {
	err error
	id  BookingID
}

type BookQueue interface {
	Add(order *InternalOrder)
	Stop()
}
