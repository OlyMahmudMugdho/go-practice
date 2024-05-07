package main

import "fmt"

func main() {

	type User struct {
		Name   string
		Roll   int
		Status bool
	}

	var me = User{"Mugdho", 14, true}
	var mila = User{"Mila", 37, true}


	fmt.Println(me)

	fmt.Printf("my info is %v \n", me)

	fmt.Printf("more formatted info about Mila %+v \n", mila);

	fmt.Printf("name of Lami is %v \n", mila.Name);
}
