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
