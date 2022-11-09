package entrypoint

import "github.com/icrxz/gointensivo/internal/order/domain"

type OrderInputDTO struct {
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func (o *OrderInputDTO) toDomain() *domain.Order {
	return &domain.Order{
		Price: o.Price,
		Tax:   o.Tax,
	}
}

func fromDomain(order *domain.Order) *OrderOutputDTO {
	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
}
