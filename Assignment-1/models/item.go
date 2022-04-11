package models

type Item struct {
	Name     string
	Price    float64
	ItemType string
}

func (o *Item) Getter() Item {
	return *o
}

func (o *Item) Setter(name string, price float64, itemType string) {
	o.Name = name
	o.Price = price
	o.ItemType = itemType
}
