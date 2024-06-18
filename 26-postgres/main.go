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

	table, err := db.Query(`
		create table books (
			id int,
			name varchar(255),
			author varchar(255)
		)
	`)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(table)
	defer db.Close()

}
