package models

type Order struct {
	I Item
	B Bill
}

func (o *Order) Getter() Order {
	return *o
}

func (o *Order) Setter(i Item, b Bill) {
	o.I = i
	o.B = b
}
