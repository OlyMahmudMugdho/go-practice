package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var factorialResults = []map[int]int{}

func main() {
	// simple example
	go greet("good morning")
	wg.Add(1)

	// example use case
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range nums {
		go DoFactorial(v)
		wg.Add(1)
	}

	// wait for the operations
	wg.Wait()

	// print the results after completion
	for _, v := range factorialResults {
		fmt.Printf("%v \n", v)
	}

}

func greet(message string) {
	defer wg.Done()
	fmt.Println(message)
}

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

func DoFactorial(num int) {
	defer wg.Done()
	var res int = factorial(num)
	factorialResults = append(factorialResults, map[int]int{num: res})
}
