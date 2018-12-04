package main

import (
	"flag"
	"fmt"
	"pizza_delivery/delivery"
	"pizza_delivery/kitchen"
	"pizza_delivery/order"
)

var orderCount = flag.Int("count", 100, "generated order count")

const (
	ovenCount = 2 //количество печей
)

func main() {
	flag.Parse()
	orders := order.GenerateOrders(*orderCount)
	for i, order := range orders {
		fmt.Printf("%v %v %v %v %v\n", i, order.Time, order.CookingDuration, order.DeliveryPoint.X, order.DeliveryPoint.Y)
	}
	kitchen := kitchen.NewKitchen(ovenCount)
	cookedOrders := kitchen.Cook(orders)
	delivery := delivery.NewDelivery(orders, cookedOrders)
	routes := delivery.CreateRoutes()
	for _, route := range routes {
		for _, order := range route {
			fmt.Printf("%v %v ", order.OrderNumber, order.DeliveryTime)
		}
		fmt.Println()
	}
}
