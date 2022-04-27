package services

import "nuclei-assignment-1/models"

func CreateBillCollection(ch chan models.Order, billList *[]models.Bill) {
	for odr := range ch {
		salestax, finalPrice := CalculateTaxAndFinalPrice(odr.ItemType, odr.ItemPrice, odr.ItemQuantity)
		b := models.Bill{
			Order:      odr,
			SalesTax:   salestax,
			FinalPrice: finalPrice,
		}
		*billList = append(*billList, b)
	}
}
