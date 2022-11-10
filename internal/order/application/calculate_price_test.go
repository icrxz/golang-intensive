package application

import (
	"testing"

	"github.com/icrxz/gointensivo/internal/order/domain"
	"github.com/stretchr/testify/assert"
)

func TestOrder_CalculateFinalPrice(t *testing.T) {
	t.Run("Should calculate final price successfully", func(t *testing.T) {
		calculatePrice := NewCalculateFinalPriceUseCase()
		order := domain.Order{ID: "123", Price: 1.1, Tax: 10}

		_, err := calculatePrice.Execute(&order)

		assert.NoError(t, err)
		assert.Equal(t, float64(11), order.FinalPrice)
	})

	t.Run("Should fails on calculate final price with invalid price", func(t *testing.T) {
		calculatePrice := NewCalculateFinalPriceUseCase()
		order := domain.Order{ID: "123", Price: 0, Tax: 10}

		_, err := calculatePrice.Execute(&order)

		assert.Error(t, err, "invalid price")
	})

	t.Run("Should fails on calculate final price with invalid tax", func(t *testing.T) {
		calculatePrice := NewCalculateFinalPriceUseCase()
		order := domain.Order{ID: "123", Price: 1.1, Tax: 0}

		_, err := calculatePrice.Execute(&order)

		assert.Error(t, err, "invalid tax")
	})
}
