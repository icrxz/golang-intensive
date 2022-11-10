package entrypoint

import (
	"encoding/json"
	"net/http"

	"github.com/icrxz/gointensivo/internal/order/application"
)

type OrderController struct {
	calculatePriceUseCase application.CalculateFinalPriceUseCase
}

func NewOrderController(calculatePriceUseCase application.CalculateFinalPriceUseCase) *OrderController {
	return &OrderController{
		calculatePriceUseCase: calculatePriceUseCase,
	}
}

func (c *OrderController) Create(w http.ResponseWriter, r *http.Request) error {
	var orderInput OrderInputDTO
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&orderInput)
	if err != nil {
		return err
	}

	order := orderInput.toDomain()
	_, err = c.calculatePriceUseCase.Execute(order)
	if err != nil {
		return err
	}

	output, err := json.Marshal(fromDomain(order))
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(output)
	return nil
}
