package application

import (
	"github.com/icrxz/gointensivo/internal/order/domain"
)

type CalculateFinalPriceUseCase struct {
	OrderRepository domain.OrderRepository
}

func NewCalculateFinalPriceUseCase(orderRepository domain.OrderRepository) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *CalculateFinalPriceUseCase) Execute(order domain.Order) (*domain.Order, error) {
	if err := order.CalculateFinalPrice(); err != nil {
		return nil, err
	}

	return &order, nil
}
