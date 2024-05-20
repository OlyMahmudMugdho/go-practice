package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
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
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")

	books = append(books, Book{"1", "Spring in Action", "Craig Walls"})

	fmt.Println("server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}

func serveHomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the home page"))
}

func getOneBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	fmt.Printf("%T\n",params["id"])
	for _, book := range books {
		if params["id"] == book.Id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	w.WriteHeader(404)
	json.NewEncoder(w).Encode(NotFound{true,"book not found"})

}


func getAllBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request)  {
	var book Book;
	json.NewDecoder(r.Body).Decode(&book);
	book.Id = fmt.Sprintf("%d", rand.Intn(100000))
	books = append(books, book)
	fmt.Printf("%T", book.Id)
	json.NewEncoder(w).Encode(books)
}