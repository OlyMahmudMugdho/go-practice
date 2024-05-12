package main

import (
	"fmt"
	"os"
)

func main() {
	var content string = "sample text 2";
	file,_ := os.Create("./some.txt");

	length,_ := file.WriteString(content);

	fmt.Println("the length of the content is ", length);

	defer file.Close();
	
}