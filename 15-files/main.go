package main

import (
	"fmt"
	"os"
)

func main() {
	var content string = "sample text 2";
	file,err := os.Create("./some.txt");

	if err != nil {
		panic(err);
	}

	length,err := file.WriteString(content);

	if err != nil {
		panic(err);
	}

	fmt.Println("the length of the content is ", length);

	defer file.Close();
	
}