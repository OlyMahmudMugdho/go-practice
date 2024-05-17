package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlParts := &url.URL{
		Scheme : "https",
		Host : "jsonplaceholder.typicode.com",
		Path : "/id",
		RawQuery : "name=Mugdho",
	}

	fullUrl := urlParts.String();
	fmt.Println(fullUrl)
}