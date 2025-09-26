package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "GoDB1"
)

var database *sql.DB

func connectToDB() {
	psqlInfo := fmt.Sprintf("host=localhost port=5432 user=postgres password=12345 dbname=GoDB1 sslmode=disable")
	var err error
	database, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database successfully connected!")
}

func createBooksTable() {
	query := `
		CREATE TABLE books (
		    id SERIAL PRIMARY KEY,
		    name VARCHAR(255) NOT NULL,
		    author VARCHAR(255) NOT NULL
		)
	`
	_, err := database.Exec(query)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Table books successfully created!")
}

func insertBook(name string, author string) {
	query :=
		`
			INSERT INTO books (name, author) VALUES ($1, $2)
		`
	_, err := database.Exec(query, name, author)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("Book successfully inserted!")
}

func insertUser1(name string, age int) {
	query := `INSERT INTO users(name, age) VALUES ($1, $2)`
	_, err := database.Exec(query, name, age)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully inserted user")
}

func getUsers1() {
	query := `SELECT * FROM users`
	rows, err := database.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id, "name:", name, "age:", age)
	}
}

func getBooks() {
	query := `SELECT * FROM books`
	rows, err := database.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("Books:")
	for rows.Next() {
		var id int
		var name string
		var author string
		err := rows.Scan(&id, &name, &author)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id, "name:", name, "author:", author)
	}
}

func main() {
	//connectToDB()
	//insertBook("test", "test")
	//getBooks()
	//createBooksTable()
	//insertBook("The lord of the rings", "J.R.R.Tolkien")
	//createTable()
	//insertUser1("John", 30)
	//getUsers1()
}
