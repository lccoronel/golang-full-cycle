package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	// create
	// database.Create(&Product{
	// 	Name:  "Mackbook Pro",
	// 	Price: 1000.00,
	// })

	// update
	// var product Product
	// database.First(&product, 1)
	// product.Name = "iPad"
	// database.Save(&product)

	//delete
	var product Product
	database.First(&product, 1)
	database.Delete(&product)
}
