package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/icrxz/gointensivo/internal/order/domain"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(ctx context.Context, order *domain.Order) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	order.ID = ID.String()

	query, err := r.Db.PrepareContext(ctx, "INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = query.ExecContext(ctx, order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}
