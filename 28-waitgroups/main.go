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

	wg.Wait()

}

func greet(message string) {
	defer wg.Done()
	fmt.Println(message)
}
