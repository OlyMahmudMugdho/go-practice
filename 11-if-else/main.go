package main

import "fmt"

func main() {
	var age = 16;
	
	if age > 18 {
		fmt.Println("You can vote")
	} else if age < 18 {
		fmt.Println("You cannot vote")
	} else {
		fmt.Println("You're 18..you can vote")
	}
}