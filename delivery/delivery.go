package delivery

import (
	"pizza_delivery/kitchen"
	"pizza_delivery/order"
)

const (
	MaxOrderCount       = 3  //Максимальное количество заказов в маршруте
	MaxDeliveryDuration = 60 //Максимальное время доставки в минутах
	DeliverySpeed       = 60 //Скорость доставки точек на минуту
)

type Delivery struct {
	orders       []order.Order
	cookedOrders kitchen.CookedOrders
}

//вычисляет матрицу периодов доставки между пунктами доставки заказов
func (d *Delivery) calcDeliveryDurationBetweenOrderPoints(cookedOrders kitchen.CookedOrders) [][]int {
	orderCount := len(cookedOrders)
	durations := make([][]int, orderCount)
	for i := 0; i < orderCount; i++ {
		durations[i] = make([]int, orderCount)
	}
	for i, cookedOrder := range cookedOrders {
		for j, cookedOrder2 := range cookedOrders[i+1:] {
			distance := getDistance(d.orders[cookedOrder.Number].DeliveryPoint, d.orders[cookedOrder2.Number].DeliveryPoint)
			duration := distance / DeliverySpeed
			durations[i][i+j+1] = duration
			durations[i+j+1][i] = duration
		}
	}
	return durations
}

//вычисляет максимальный запас времени на доставку после приготовления последнего заказа из списка
func (d *Delivery) getMaxDeliveryTime(cookedOrders kitchen.CookedOrders) []int {
	cookedOrderCount := len(cookedOrders)
	times := make([]int, cookedOrderCount)
	lastOrderCookedTime := cookedOrders[cookedOrderCount-1].Time //время приготовления последнего заказа
	for i, order := range cookedOrders {
		times[i] = MaxDeliveryDuration - (lastOrderCookedTime - d.orders[order.Number].Time)
	}
	return times
}

//инициирует маршруты из начальной точки
func (d *Delivery) calcStartRawRoutes(cookedOrders kitchen.CookedOrders, maxDeliveryTime []int) []rawRoute {
	routes := make([]rawRoute, 0)
	for i, cookedOrder := range cookedOrders {
		deliveryDuration := getDistance(order.Point{}, d.orders[cookedOrder.Number].DeliveryPoint) / DeliverySpeed
		if deliveryDuration <= maxDeliveryTime[i] {
			r := createRawRoute(len(cookedOrders))
			r.AddOrder(i, deliveryDuration)
			routes = append(routes, r)
		}
	}
	return routes
}

//формирует маршрут с наименьшей общей продолжительностью по доставке с учетом запаса времени
func (d *Delivery) calcFitRoute(cookedOrders kitchen.CookedOrders) (bool, Route) {
	deliveryDuration := d.calcDeliveryDurationBetweenOrderPoints(cookedOrders)
	maxDeliveryTime := d.getMaxDeliveryTime(cookedOrders)
	cookedOrderCount := len(cookedOrders)
	possibleRoutes := d.calcStartRawRoutes(cookedOrders, maxDeliveryTime)
	route := Route{}
	routeDuration := MaxDeliveryDuration + 1
	for len(possibleRoutes) > 0 {
		rawRoute := possibleRoutes[0]                          //берем первую цепочку
		possibleRoutes = possibleRoutes[1:]                    //удаляем первую цепочку
		current, currentDuration, _ := rawRoute.GetLastOrder() //берем последний заказ из цепочки
		for i := 0; i < cookedOrderCount; i++ {                //проходим по всем остальным заказам
			if !rawRoute.IsAddedOrder(i) { //если по этому заказу не проходили в этой цепочке и не текущий заказ
				duration := deliveryDuration[current][i] //берем период доставки от текущего заказа до следующего
				if currentDuration+duration < maxDeliveryTime[i] { //если период доставки подходит
					nextRawRoute := rawRoute.Copy()
					nextRawRoute.AddOrder(i, duration) //добавляем заказ в цепочку
					_, duration, length := nextRawRoute.GetLastOrder()
					if length == cookedOrderCount { // построен ли маршрут
						if duration < routeDuration { //минимален ли по периоду
							routeDuration = duration
							route = nextRawRoute.GetFullRoute(cookedOrders) //строим полный маршрут с временем доставки
						}
					} else {
						possibleRoutes = append(possibleRoutes, nextRawRoute)
					}
				}
			}
		}
	}
	return len(route) > 0, route
}

//выбираем заказы поступившие ранее времени приготовления заданного элемента
func (d *Delivery) getFitOrdersByTime(i int) kitchen.CookedOrders {
	maxOrderCount := len(d.cookedOrders) - i
	if maxOrderCount > MaxOrderCount {
		maxOrderCount = MaxOrderCount
	}
	for j := i + 1; j < i+maxOrderCount; j++ {
		if d.cookedOrders[i].Time < d.orders[d.cookedOrders[j].Number].Time {
			return d.cookedOrders[i:j]
		}
	}
	return d.cookedOrders[i : i+maxOrderCount]
}

//вычисляет максимальный маршрут с заданного элемента
func (d *Delivery) calcMaxRoute(i int) (Route, int) {
	cookedOrders := d.getFitOrdersByTime(i)
	for len(cookedOrders) > 0 {
		if len(cookedOrders) == 1 { //заказ не может быть сгруппирован
			cookedOrder := cookedOrders[0]
			deliveryDuration := getDistance(order.Point{X: 0, Y:0}, d.orders[cookedOrder.Number].DeliveryPoint) / DeliverySpeed
			return Route{
				OrderDelivery{
					OrderNumber:  cookedOrder.Number,
					DeliveryTime: cookedOrder.Time + deliveryDuration,
				},
			}, 1
		}
		if is, route := d.calcFitRoute(cookedOrders); is == true { //вычисление возможности маршрута сгруппированных заказов
			return route, len(route)
		}
		cookedOrders = cookedOrders[:len(cookedOrders)-1] //убирается последний заказ из списка
	}
	return Route{}, 1
}

func (d *Delivery) CreateRoutes() []Route {
	routes := make([]Route, 0)
	for i := 0; i < len(d.orders); {
		route, count := d.calcMaxRoute(i)
		i += count
		routes = append(routes, route)
	}
	return routes
}

func NewDelivery(orders []order.Order, cookedOrders kitchen.CookedOrders) *Delivery {
	return &Delivery{
		orders:       orders,
		cookedOrders: cookedOrders,
	}
}
