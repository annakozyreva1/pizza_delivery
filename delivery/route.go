package delivery

import (
	"pizza_delivery/kitchen"
)

type OrderDelivery struct {
	OrderNumber  int //номер заказа из списка
	DeliveryTime int //время доставки в минутах
}

type Route []OrderDelivery //маршрут отправки заказов

//цепочка возможного подходящего маршрута
type rawRoute struct {
	orderOrderliness      []int //значения порядка в маршруте элементов подсписка заказов
	orderDeliveryDuration []int //значения продолжительности доставки в маршруте элементов подсписка заказов
	lastOrderIndex        int   //номер последнего пройденого элемента из подсписка заказов
	length                int   //количество элементов в маршруте
}

func (r *rawRoute) AddOrder(i int, duration int) {
	r.orderOrderliness[i] = r.length
	r.orderDeliveryDuration[i] = r.orderDeliveryDuration[r.lastOrderIndex] + duration
	r.lastOrderIndex = i
	r.length = r.length + 1
}

func (r *rawRoute) GetLastOrder() (i int, duration int, length int) {
	return r.lastOrderIndex, r.orderDeliveryDuration[r.lastOrderIndex], r.length
}

func (r *rawRoute) GetFullRoute(cookedOrders kitchen.CookedOrders) Route {
	route := make(Route, r.length)
	lastOrderCookTime := cookedOrders[len(cookedOrders)-1].Time
	for i := 0; i < r.length; i++ {
		route[r.orderOrderliness[i]] = OrderDelivery{
			OrderNumber:  cookedOrders[i].Number,
			DeliveryTime: lastOrderCookTime + r.orderDeliveryDuration[i],
		}
	}
	return route
}

func (r *rawRoute) IsAddedOrder(i int) bool {
	return r.orderOrderliness[i] > 0 || r.lastOrderIndex == i
}

func (r *rawRoute) Copy() rawRoute {
	route := createRawRoute(len(r.orderDeliveryDuration))
	copy(route.orderDeliveryDuration, r.orderDeliveryDuration)
	copy(route.orderOrderliness, r.orderOrderliness)
	route.length = r.length
	route.lastOrderIndex = r.lastOrderIndex
	return route
}

func createRawRoute(orderCount int) rawRoute {
	return rawRoute{
		orderDeliveryDuration: make([]int, orderCount),
		orderOrderliness:      make([]int, orderCount),
	}
}
