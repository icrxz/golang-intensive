package repository

import (
	"context"
	"database/sql"

	"github.com/icrxz/gointensivo/internal/order/domain"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) domain.OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(ctx context.Context, order *domain.Order) error {
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

func (r *OrderRepository) GetTotal(ctx context.Context) (int, error) {
	var total int

	err := r.Db.QueryRowContext(ctx, "SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
