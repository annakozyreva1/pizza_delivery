package order

import (
	"math"
	"math/rand"
	"time"
)

var (
	OrdersDurationRange  = [2]int{1, 30}       //диапазон возможных значений между заказами
	CookingDurationRange = [2]int{10, 30}      //диапазон возможных значений периода приготовленяия
	PointRange           = [2]int{-1000, 1000} //область возможных пунктов доставки
)

func generatePoisson(lambda float64) int {
	L := math.Pow(math.E, -lambda)
	k := 0
	var p float64 = 1.0
	for p > L {
		k++
		p *= rand.Float64()
	}
	return k - 1
}

func generateInRange(Range [2]int) int {
	return Range[0] + rand.Intn(Range[1]-Range[0]+1)
}

func getTimePeriod(timeRange [2]int) int {
	return generateInRange(timeRange)
}

func getTimePeriodByPoissonDistribution(timeRange [2]int) int {
	v := generatePoisson(float64(timeRange[1]-timeRange[0]) / 2.0)
	if v > timeRange[1] {
		v = timeRange[1]
	}
	return v
}

func getPoint(pointRange [2]int) Point {
	return Point{
		X: generateInRange(pointRange),
		Y: generateInRange(pointRange),
	}
}

func GenerateOrders(count int) []Order {
	rand.Seed(time.Now().UTC().UnixNano())
	orders := make([]Order, count)
	nextOrderTime := 0
	for i, _ := range orders {
		orders[i] = Order{
			Time:            nextOrderTime,
			CookingDuration: getTimePeriod(CookingDurationRange),
			DeliveryPoint:   getPoint(PointRange),
		}
		nextOrderTime += getTimePeriodByPoissonDistribution(OrdersDurationRange)
	}
	return orders
}
