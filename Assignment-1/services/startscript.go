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
		var (
			i models.Item
			b models.Bill
			o models.Order
		)
		i.Setter(iname, iprice, itype)
		b.Setter(i, quantity)
		o.Setter(i, b)
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
