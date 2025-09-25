package main

import (
	"assignment-2/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=12345 dbname=GoDB2 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	return db
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Table 'users' migrated!")
}

func insertUser(db *gorm.DB, name string, age int) {
	user := model.User{Name: name, Age: age}
	db.Create(&user)
	fmt.Printf("Inserted: %s (%d)\n", name, age)
}

func findAll(db *gorm.DB) {
	var users []model.User
	db.Find(&users)
	fmt.Println("Users:")
	for _, u := range users {
		fmt.Printf("ID=%d | Name=%s | Age=%d\n", u.ID, u.Name, u.Age)
	}
}

func main() {
	db := connectToDatabase()
	findAll(db)
	//insertUser(db, "Vasya", 20)
}
