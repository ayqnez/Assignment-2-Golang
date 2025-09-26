package main

import (
	"assignment-2/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dataBase *gorm.DB

func connectToDatabase() {
	dsn := "host=localhost user=postgres password=12345 dbname=GoDB2 port=5432 sslmode=disable"
	var err error
	dataBase, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
}

func migrate() {
	err := dataBase.AutoMigrate(&model.Product{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migrated products")
}

func insertUser(db *gorm.DB, name string, age int) {
	user := model.User{Name: name, Age: age}
	db.Create(&user)
	fmt.Printf("Inserted: %s (%d)\n", name, age)
}

func insertProduct(name string, price int) {
	product := model.Product{Name: name, Price: price}
	dataBase.Create(&product)
	fmt.Printf("Inserted product", name, price)
}

func findAll(db *gorm.DB) {
	var users []model.User
	db.Find(&users)
	fmt.Println("Users:")
	for _, u := range users {
		fmt.Printf("ID=%d | Name=%s | Age=%d\n", u.ID, u.Name, u.Age)
	}
}

func findAllProducts() {
	var books []model.Product
	dataBase.Find(&books)
	fmt.Println("Products:")
	for _, book := range books {
		fmt.Printf("ID: %d | Name=%s | Price=%d\n", book.ID, book.Name, book.Price)
	}
}

func main() {
	connectToDatabase()
	//migrate()
	//insertProduct("12345", 12345)
	findAllProducts()
}
