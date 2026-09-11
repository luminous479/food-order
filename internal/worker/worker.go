package worker

import (
	"fmt"
	"sync"

	"github.com/luminous479/food-order/internal/order"
)

func Worker(orders <-chan order.Order, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orders {
		fmt.Printf("Processing order ID: %d, Customer: %s, Food: %s\n", order.ID, order.Customer, order.Food)

	}
}
