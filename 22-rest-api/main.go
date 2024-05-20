package main

import (
	"encoding/json"
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

type NotFound struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

var books []Book

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", serveHomePage).Methods("GET")
	router.HandleFunc("/books/{id}", getOneBook).Methods("GET")

	books = append(books, Book{"1", "Spring in Action", "Craig Walls"})

	fmt.Println("server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}

func serveHomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the home page"))
}

func getOneBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	for _, book := range books {
		if params["id"] == book.Id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	w.WriteHeader(404)
	json.NewEncoder(w).Encode(NotFound{true,"book not found"})

	return
}
