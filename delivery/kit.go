package delivery

import (
	"math"
	"pizza_delivery/order"
)

func getDistance(p1 order.Point, p2 order.Point) int {
	x := (p2.X - p1.X) * (p2.X - p1.X)
	y := (p2.Y - p1.Y) * (p2.Y - p1.Y)
	d := math.Sqrt(float64(x + y))
	return int(d)
}
