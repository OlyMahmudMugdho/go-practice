package main

import "fmt"

func main() {
	go SumOfTen()
	fmt.Println("hello")
}

func SumOfTen() int8 {
	var sum int8 = 0

	for i := 1; i <= 10; i++ {
		sum = sum + int8(i)
	}
	return sum
}
