package services

import (
	"nuclei-assignment-1/models"
	"nuclei-assignment-1/services/src"
)

func StartScript() {
	/*
		[✅]take input
		[✅]validate input
		[✅]store in order struct (o1)
		[✅]create a bill for the order (b1)
		**
		[✅]display the bill (b1)

		**[✅]take another input (o2,o3...)
		goto step 3
	*/
	var cart []models.Order
	for {
	start:
		iname, itype, iprice, quantity := src.ReadInput()
		tax, finalPrice := CalculateTaxAndFinalPrice(itype, iprice, quantity)
		o := models.Order{
			ItemName:     iname,
			ItemType:     itype,
			ItemPrice:    iprice,
			ItemQuantity: quantity,
			SalesTax:     tax,
			FinalPrice:   finalPrice,
		}
		cart = append(cart, o)
		ok := src.MoreOrders()
		if ok {
			goto start
		}
		for _, order := range cart {
			src.DisplayBill(order.Getter())
		}
		break
	}
}
