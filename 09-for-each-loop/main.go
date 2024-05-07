package main

import "fmt"

func main() {
	students := make(map[int]string);

	students[0] = "Mila"
	students[1] = "Mugdho"
	students[2] = "Meghla"

	for key,value := range students {
		fmt.Printf("roll : %v name : %v \n", key + 1, value);
	}
	
}