package application

import (
	"context"

	"github.com/icrxz/gointensivo/internal/order/domain"
)

type OrderInputDTO struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository domain.OrderRepository
}

func NewCalculateFinalPriceUseCase(orderRepository domain.OrderRepository) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *CalculateFinalPriceUseCase) Execute(ctx context.Context, input OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := domain.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	order.FinalPrice = order.Price * order.Tax
	if err := order.IsValid(); err != nil {
		return nil, err
	}

	if err = c.OrderRepository.Save(ctx, order); err != nil {
		return nil, err
	}

	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
