package order

type Point struct {
	X int
	Y int
}

type Order struct {
	Time            int   //время поступления заказа в минутах
	CookingDuration int   //время приготовления в минутах
	DeliveryPoint   Point //координата доставки
}
