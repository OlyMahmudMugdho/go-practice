package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id     string `json:book_id`
	Name   string `json:name`
	Author string `json:author`
}

func main() {
	router := mux.NewRouter();
	router.HandleFunc("/", serveHomePage).Methods("GET")
	fmt.Println("server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
	
}

func serveHomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the home page"))
}
