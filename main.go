package main

import (
	"sync"

	ord "github.com/luminous479/food-order/internal/order"
	"github.com/luminous479/food-order/internal/worker"
)

func main() {
   
	var wg sync.WaitGroup

	orders := make(chan ord.Order)
    wg.Add(2)

	go func(){
		defer close(orders)
		orders <-ord.Order{ID: 1, Customer: "John Doe", Food: "Pizza"}
		orders <-ord.Order{ID: 2, Customer: "Jane Smith", Food: "Burger"}
		orders <-ord.Order{ID: 3, Customer: "Alice Johnson", Food: "Sushi"}
		orders <-ord.Order{ID: 4, Customer: "Bob Brown", Food: "Pasta"}
		orders <-ord.Order{ID: 5, Customer: "Charlie Davis", Food: "Salad"}
	}()

	go worker.Worker(orders, &wg)
	go worker.Worker(orders, &wg)

	wg.Wait()

	
}
