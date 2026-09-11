package main

import (
	"github.com/luminous479/food-order/internal/order"
	"github.com/luminous479/food-order/internal/worker"
)

func main() {

	order := make(chan order.Order)

	go worker.Worker(order)

	g
}
