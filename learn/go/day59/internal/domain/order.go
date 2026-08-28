package domain

import "errors"

var ErrInvalidOrder = errors.New("invalid order")

type Order struct {
	ID        int64
	Customer  string
	Item      string
	Qty       int
	UnitPrice float64
	Total     float64
}

func (o Order) Validate() error {
	if o.Customer == "" || o.Item == "" || o.Qty <= 0 || o.UnitPrice <= 0 {
		return ErrInvalidOrder
	}
	return nil
}

func (o Order) ComputeTotal() float64 {
	return float64(o.Qty) * o.UnitPrice
}
