package main

import (
	"fmt"
	"os"
)

func main() {
	var content string = "sample text";
	file,err := os.Create("./some.txt");

	if err != nil {
		panic(err);
	}

	length,err := file.WriteString(content);

	if err != nil {
		panic(err);
	}

	fmt.Println("the length of the content is ", length);


	readFile()

	defer file.Close();
	
}


func readFile() {
	dataBytes,_ := os.ReadFile("./some.txt");
	fmt.Println("The data bytes in the file : ");
	fmt.Println(dataBytes);

	fmt.Println("The text in the file : ");
	fmt.Println(string(dataBytes));
}