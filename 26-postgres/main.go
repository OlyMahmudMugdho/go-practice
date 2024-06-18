package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("hello")
	connStr := "user=postgres dbname=postgres password=mysecretpassword sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := createTable(db)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)

	/* rows, err := droptTable(db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)
	// for dropping table

	*/

	defer db.Close()

}

func droptTable(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query(`
	  	drop table books
	`)
	return rows, err
}

func createTable(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query(`
		create table books (
			id int,
			name varchar(255),
			author varchar(255)
		)
	`)
	return rows, err
}
