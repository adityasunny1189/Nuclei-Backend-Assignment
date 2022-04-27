package services

import (
	"database/sql"
	"fmt"
	"nuclei-assignment-1/models"
)

const (
	queryStatement = "SELECT * FROM cart"
)

func FetchDataFromDB(ch chan models.Order, db *sql.DB, orderlist *[]models.Order) {
	rows, qerr := db.Query(queryStatement)
	if qerr != nil {
		fmt.Println(qerr)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var odr models.Order
		if err := rows.Scan(&odr.ItemName, &odr.ItemPrice, &odr.ItemType, &odr.ItemQuantity); err != nil {
			fmt.Println(err)
			return
		}
		// odr.DisplayOrder()
		*orderlist = append(*orderlist, odr)
		ch <- odr
	}
}
