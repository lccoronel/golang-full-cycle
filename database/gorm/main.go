package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

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
	// var product Product
	// database.First(&product, 1)
	// database.Delete(&product)

	// belongs to
	// category := Category{Name: "Eletornicos"}
	// database.Create(&category)
	// database.Create(&Product{
	// 	Name:       "Notebook",
	// 	Price:      1000.00,
	// 	CategoryID: category.ID,
	// })
	// database.Create(&SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: 1,
	// })

	// find all
	// var products []Product
	// database.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	var categories []Category
	err = database.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")

		for _, product := range category.Products {
			println("- ", product.Name, "Serial Number:", product.SerialNumber.Number)
		}
	}
}
