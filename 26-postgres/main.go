package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("hello")
	connStr := "user=postgres dbname=postgres password=mysecretpassword"
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
