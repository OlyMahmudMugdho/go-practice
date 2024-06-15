package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	port := "8080"
	fmt.Printf("server is running on port %v\n", port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}
