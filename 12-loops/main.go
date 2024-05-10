package main

import "fmt"

func main() {

	// basic for loop

	/*
		for i := 0; i < 10; i++ {
			fmt.Println(i);
		}
	*/

	var num int = 10;

	for num >= 0 {
		fmt.Println(num);
		num--;
	}
	
}