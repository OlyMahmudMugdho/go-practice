package main

import "fmt";

func main() {
	sayHello();
	Message("mugdho")
}

func sayHello()  {
	fmt.Println("Hello");
}

func Message(msg string)  {
	fmt.Println("a message to ", msg)
}