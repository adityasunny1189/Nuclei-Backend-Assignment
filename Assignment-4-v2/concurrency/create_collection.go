package concurrency

import (
	"gorm-orm-connection/models"
	"gorm-orm-connection/services"
)

func CreateCollection(ch <-chan models.Order, ch2 chan<- models.Bill, collection2 []models.Bill) {
	for val := range ch {
		st, fp := services.CalculateTaxAndFinalPrice(val)
		b := models.Bill{
			Order:      val,
			SalesTax:   st,
			FinalPrice: fp,
		}
		collection2 = append(collection2, b)
		ch2 <- b
	}
	close(ch2)
}
