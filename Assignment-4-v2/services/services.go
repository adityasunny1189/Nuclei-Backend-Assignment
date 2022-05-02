package services

import (
	"fmt"
	"gorm-orm-connection/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitiateDB() *gorm.DB {
	dsn := "root:Adisunny123@tcp(127.0.0.1:3306)/cart?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("connected")
	db.AutoMigrate(&models.Order{})
	return db
}

func FillOneMillionDataInJSON(db *gorm.DB) {
	for i := 0; i < 200000; i++ {
		o := models.Order{
			ItemName:     "random name",
			ItemType:     "raw",
			ItemPrice:    1000,
			ItemQuantity: 4,
		}
		db.Create(&o)
	}
}
