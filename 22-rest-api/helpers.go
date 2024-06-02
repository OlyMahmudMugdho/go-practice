package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.URL)

	msg := make(map[string]string);

	msg["message"] = "hello";

	messageJson,_ := json.Marshal(msg);

	w.Write(messageJson)
}