package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	t.Run("Should create an order successfully with all params", func(t *testing.T) {
		order, err := NewOrder("123", 1.1, 1)

		assert.NoError(t, err)
		assert.Equal(t, "123", order.ID)
		assert.Equal(t, float64(1.1), order.Price)
		assert.Equal(t, float64(1), order.Tax)
	})

	t.Run("Should return an error when create an order with invalid price", func(t *testing.T) {
		order, err := NewOrder("123", 0, 1)

		assert.Nil(t, order)
		assert.Error(t, err, "invalid price")
	})

	t.Run("Should return an error when create an order with invalid price", func(t *testing.T) {
		order, err := NewOrder("123", 1.1, 0)

		assert.Nil(t, order)
		assert.Error(t, err, "invalid tax")
	})
}

func TestOrder_IsValid(t *testing.T) {
	t.Run("Should fail with an empty ID", func(t *testing.T) {
		order := Order{}

		assert.Error(t, order.IsValid(), "invalid ID")
	})

	t.Run("Should fail with an invalid price", func(t *testing.T) {
		order := Order{ID: "123"}

		assert.Error(t, order.IsValid(), "invalid price")
	})

	t.Run("Should fail with an invalid tax", func(t *testing.T) {
		order := Order{ID: "123", Price: 1.11}

		assert.Error(t, order.IsValid(), "invalid tax")
	})

	t.Run("Should validate an order successfully with all params", func(t *testing.T) {
		order := Order{ID: "123", Price: 1.11, Tax: 1}

		assert.NoError(t, order.IsValid())
	})
}

func TestOrder_CalculateFinalPrice(t *testing.T) {
	t.Run("Should calculate final price successfully", func(t *testing.T) {
		order := Order{ID: "123", Price: 1.1, Tax: 10}

		err := order.CalculateFinalPrice()

		assert.NoError(t, err)
		assert.Equal(t, float64(11), order.FinalPrice)
	})

	t.Run("Should fails on calculate final price with invalid price", func(t *testing.T) {
		order := Order{ID: "123", Price: 0, Tax: 10}

		err := order.CalculateFinalPrice()

		assert.Error(t, err, "invalid price")
	})

	t.Run("Should fails on calculate final price with invalid tax", func(t *testing.T) {
		order := Order{ID: "123", Price: 1.1, Tax: 0}

		err := order.CalculateFinalPrice()

		assert.Error(t, err, "invalid tax")
	})
}
