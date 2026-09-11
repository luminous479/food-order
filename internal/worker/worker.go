package worker

import (
	"fmt"

	"github.com/luminous479/food-order/internal/order"
)


func Worker(orders <-chan order.Order){



	 for order := range orders {

		fmt.Printf("Processing order ID: %d, Customer: %s, Food: %s\n", order.ID, order.Customer, order.Food)
		
	 }
}