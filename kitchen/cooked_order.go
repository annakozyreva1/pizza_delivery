package kitchen

type CookedOrder struct {
	Number int //номер заказа из списка
	Time   int //время приготовления заказа
}

type CookedOrders []CookedOrder //список приготовленных заказов

func (c CookedOrders) Len() int {
	return len(c)
}

func (c CookedOrders) Less(i, j int) bool {
	return c[i].Time < c[j].Time
}

func (c CookedOrders) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
