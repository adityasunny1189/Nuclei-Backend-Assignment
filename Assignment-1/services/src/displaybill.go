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

func DisplayBill(o models.Order) {
	fmt.Printf(DisplayBillContent, o.B.ItemName, o.B.ItemPrice, o.B.ItemQuantity, o.B.SalesTax, o.B.FinalPrice)

}
