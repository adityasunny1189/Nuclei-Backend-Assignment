package models

type Order struct {
	ItemName     string
	ItemPrice    float64
	ItemType     string
	SalesTax     float64
	FinalPrice   float64
	ItemQuantity int
}
