package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{}, &Category{})

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

	// find all
	var products []Product
	database.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
