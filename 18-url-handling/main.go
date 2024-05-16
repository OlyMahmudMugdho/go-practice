package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://jsonplaceholder.typicode.com:8080/important?completed=true";

func main() {
	fmt.Printf("main url : %v \n", myUrl);
	parsedUrl, _ := url.Parse(myUrl);
	fmt.Println("protocol" + parsedUrl.Scheme)
	fmt.Println("host : " + parsedUrl.Host)
	fmt.Println("path : " + parsedUrl.Path)
	fmt.Println("query : " + parsedUrl.RawQuery)
	fmt.Println("port : " + parsedUrl.Port())

	fmt.Printf("Query is a map, :  ")
	fmt.Println(parsedUrl.Query())
	
}