package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	go greet("good morning")
	wg.Add(1)

	wg.Wait()
}

func greet(message string) {
	defer wg.Done()
	fmt.Println(message)
}

func factorial(num int) int {
	defer wg.Done()
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorial(num-1)
}
