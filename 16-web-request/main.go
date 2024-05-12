package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://jsonplaceholder.typicode.com"

func main() {
	response, _ := http.Get(url);
	responseText,_ := io.ReadAll(response.Body)
	htmlData := string(responseText)
	fmt.Println(htmlData)


	file,_ := os.Create("index.html");

	file.WriteString(htmlData);
}