package app

import (
	"database/sql"
	"fmt"
	"nuclei-assignment-1/models"
	"nuclei-assignment-1/services"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func StartApp() {
	/*
		-> create a struct order and bill
		-> create two slice of type order and bill
		-> create a DB connection
		-> store the order value in DB
		-> concurrently fetch order details from DB and store it in orders and at the same time
		also calculate the bill for the order in orders and store it in bills
	*/
	var db *sql.DB
	var orders []models.Order
	var bills []models.Bill
	ch := make(chan models.Order)
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "mystore",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected to the database mystore")

	go services.CreateBillCollection(ch, &bills)
	services.FetchDataFromDB(ch, db, &orders)

	time.Sleep(time.Second * 2)

	for _, bill := range bills {
		bill.DisplayBill()
	}
}
