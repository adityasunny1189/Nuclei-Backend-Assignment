package concurrency

import "gorm-orm-connection/models"

func FetchData(ch chan<- models.Order, odrs, collection1 []models.Order) {
	for _, odr := range odrs {
		collection1 = append(collection1, odr)
		ch <- odr
	}
	close(ch)
}
