package main

import "fmt"


func main()  {
	var num int = 23;
	var address *int = &num;

	fmt.Println(num);
	fmt.Print("address of number is ")
	fmt.Println(address);

	fmt.Println("the value stored in the address is ", *address);
}