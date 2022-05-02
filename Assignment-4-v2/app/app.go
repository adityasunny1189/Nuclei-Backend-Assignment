package app

import (
	"gorm-orm-connection/concurrency"
	"gorm-orm-connection/models"
	"gorm-orm-connection/services"
)

func StartApp() {
	db := services.InitiateDB()
	var cart []models.Order
	var collection1 []models.Order
	var collection2 []models.Bill
	ch1 := make(chan models.Order)
	ch2 := make(chan models.Bill)
	db.Find(&cart)

	go concurrency.FetchData(ch1, cart, collection1)
	go concurrency.CreateCollection(ch1, ch2, collection2)

	for b := range ch2 {
		b.DisplayBill()
	}

	// for filling 1 million data in json file
	// services.FillOneMillionDataInJSON()
}
