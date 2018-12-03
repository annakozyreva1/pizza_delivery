package kitchen

import (
	"pizza_delivery/order"
	"sort"
)

type Kitchen struct {
	ovensFreeTimes []int
}

func (k *Kitchen) cookOrder(orderTime, cookingTime int) int {
	sort.Ints(k.ovensFreeTimes)          // отсортировать время освобождения печей по мнимальному
	if orderTime > k.ovensFreeTimes[0] { // если время заказа больше минимального времени освобождения печи то устанавливаем время заказа
		k.ovensFreeTimes[0] = orderTime
	}
	k.ovensFreeTimes[0] += cookingTime //добавляем к времени наиболее быстроосвободившейся печи время приготовления
	return k.ovensFreeTimes[0]
}

func (k *Kitchen) Cook(orders []order.Order) []CookedOrder {
	cookedOrders := make(CookedOrders, len(orders))
	for i, order := range orders {
		cookedOrders[i] = CookedOrder{
			Number: i,
			Time:   k.cookOrder(order.Time, order.CookingDuration),
		}
	}
	sort.Sort(cookedOrders)
	return cookedOrders
}

func NewKitchen(ovenCount int) *Kitchen {
	return &Kitchen{
		ovensFreeTimes: make([]int, ovenCount, ovenCount),
	}
}
