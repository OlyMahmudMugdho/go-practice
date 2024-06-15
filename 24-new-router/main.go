package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	port := "8080"


	router.HandleFunc("GET /", sayHello)
	


	fmt.Printf("server is running on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type","text/html");
	json.NewEncoder(w).Encode("hello")
}
