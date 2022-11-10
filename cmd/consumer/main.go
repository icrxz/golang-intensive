package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/icrxz/gointensivo/internal/order/application"
	"github.com/icrxz/gointensivo/internal/order/infra/repository"
	"github.com/icrxz/gointensivo/internal/order/infra/repository/database"
	"github.com/icrxz/gointensivo/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ctx := context.Background()

	db := database.Load()
	defer db.Close()

	orderRepository := repository.NewOrderRepository(db)

	calculateFinalPrice := application.NewCalculateFinalPriceUseCase(orderRepository)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consumer(ch, msgs)

	for msg := range msgs {
		var inputDTO application.OrderInputDTO
		if err := json.Unmarshal(msg.Body, &inputDTO); err != nil {
			panic(err)
		}

		output, err := calculateFinalPrice.Execute(ctx, inputDTO)
		if err != nil {
			fmt.Println(err.Error())
		}

		msg.Ack(false)
		fmt.Printf("Output: %v\n", output)
	}
}
