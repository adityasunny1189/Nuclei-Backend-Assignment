package src

import (
	"fmt"
	"nuclei-assignment-1/models"
)

func DisplayBill(o models.Order) {
	fmt.Printf("\n-----------BILL DETAILS-----------\n\n")
	fmt.Printf("ITEM NAME: %s\n", o.B.ItemName)
	fmt.Printf("ITEM PRICE: %f\n", o.B.ItemPrice)
	fmt.Printf("ITEM QUANTITY: %d\n", o.B.ItemQuantity)
	fmt.Printf("SALES TAX: %.3f\n", o.B.SalesTax)
	fmt.Printf("FINAL PRICE: %.4f\n", o.B.FinalPrice)
	fmt.Printf("\n----------HAPPY SHOOPING----------\n\n")
}
