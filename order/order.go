package order

type Point struct {
	X int //x координата точки доставки заказа
	Y int //y координата точки доставки заказа
}

type Order struct {
	Time            int   //время поступления заказа в минутах
	CookingDuration int   //время приготовления в минутах
	DeliveryPoint   Point //координата доставки
}
