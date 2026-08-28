package service

import (
	"sync"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day59/internal/domain"
)

type OrderService struct {
	mu     sync.Mutex
	nextID int64
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) Place(order domain.Order) (domain.Order, error) {
	if err := order.Validate(); err != nil {
		return domain.Order{}, err
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nextID++
	order.ID = s.nextID
	order.Total = order.ComputeTotal()
	return order, nil
}
