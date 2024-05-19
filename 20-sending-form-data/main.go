package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const serverUrlPost string = "http://localhost:3000/post-end"
func main() {
	data := url.Values{};

	data.Add("name", "Mugdho");
	data.Add("roll", "210614");

	response,err := http.PostForm(serverUrlPost, data);

	if err != nil {
		panic(err)
	}

	// content,_ := ioutil.ReadAll(response.Body); 
	// NB  : ioutil.ReadAll() is deprecated

	content,_ := io.ReadAll(response.Body);
	fmt.Println(string(content))

	fmt.Println("some small change")
}