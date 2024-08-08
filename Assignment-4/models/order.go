package models

import "fmt"

const (
	DisplayOrderContent = `-----------BILL DETAILS-----------
ITEM NAME: %s
ITEM PRICE: %f
ITEM QUANTITY: %d
----------HAPPY SHOOPING----------
`
	DisplayBillContent = `-----------BILL DETAILS-----------
ITEM NAME: %s
ITEM PRICE: %f
ITEM QUANTITY: %d
SALES TAX: %.3f
FINAL PRICE: %.4f
----------HAPPY SHOOPING----------
`
)

type Order struct {
	ItemName     string
	ItemPrice    float64
	ItemType     string
	ItemQuantity int
}

type Bill struct {
	Order
	SalesTax   float64
	FinalPrice float64
}

func (b *Bill) DisplayBill() {
	fmt.Printf(DisplayBillContent, b.ItemName, b.ItemPrice, b.ItemQuantity, b.SalesTax, b.FinalPrice)
}

func (o *Order) DisplayOrder() {
	fmt.Printf(DisplayOrderContent, o.ItemName, o.ItemPrice, o.ItemQuantity)
}
