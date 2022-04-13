package src

import (
	"fmt"
	"nuclei-assignment-1/models"
)

const (
	DisplayBillContent = `-----------BILL DETAILS-----------
ITEM NAME: %s
ITEM PRICE: %f
ITEM QUANTITY: %d
SALES TAX: %.3f
FINAL PRICE: %.4f
----------HAPPY SHOOPING----------
`
)

func DisplayBill(odr models.Order) {
	fmt.Printf(DisplayBillContent, odr.ItemName, odr.ItemPrice, odr.ItemQuantity, odr.SalesTax, odr.FinalPrice)
}
