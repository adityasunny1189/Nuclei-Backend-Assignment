package src

import (
	"fmt"
	"nuclei-assignment-1/models"
)

func DisplayBill(o models.Order) {
	fmt.Printf(`-----------BILL DETAILS-----------
ITEM NAME: %s
ITEM PRICE: %f
ITEM QUANTITY: %d
SALES TAX: %.3f
FINAL PRICE: %.4f
----------HAPPY SHOOPING----------
`, o.B.ItemName, o.B.ItemPrice, o.B.ItemQuantity, o.B.SalesTax, o.B.FinalPrice)

}
