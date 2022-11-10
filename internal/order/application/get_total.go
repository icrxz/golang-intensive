package application

import (
	"context"

	"github.com/icrxz/gointensivo/internal/order/domain"
)

type GetTotalOutputDTO struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository domain.OrderRepository
}

func NewGetTotalUseCase(orderRepository domain.OrderRepository) *GetTotalUseCase {
	return &GetTotalUseCase{
		OrderRepository: orderRepository,
	}
}

func (g *GetTotalUseCase) Execute(ctx context.Context) (*GetTotalOutputDTO, error) {
	total, err := g.OrderRepository.GetTotal(ctx)
	if err != nil {
		return nil, err
	}

	return &GetTotalOutputDTO{
		Total: total,
	}, nil
}
