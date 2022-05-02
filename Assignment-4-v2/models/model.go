package models

import "fmt"

type Order struct {
	ID           int     `gorm:"primaryKey"`
	ItemName     string  `json:"item_name"`
	ItemPrice    float64 `json:"item_price"`
	ItemType     string  `json:"item_type"`
	ItemQuantity int     `json:"item_quantity"`
}

type Bill struct {
	Order
	SalesTax   float64
	FinalPrice float64
}

func (b *Bill) DisplayBill() {
	fmt.Printf("Id: %d Item name : %s, Item Price: %f, Sales Tax: %f, Final Price: %f\n",
		b.ID, b.ItemName, b.ItemPrice, b.SalesTax, b.FinalPrice)
}
