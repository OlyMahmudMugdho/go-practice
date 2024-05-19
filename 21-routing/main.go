package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("working...")
	r := mux.NewRouter()
	r.HandleFunc("/", sayHello).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("running server")
}
