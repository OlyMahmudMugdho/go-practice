package main

import (
	"fmt"
	"time"
)

func main() {
	go greet("hello")
	greet("world")
}

func greet(message string) {
	start := time.Now().UnixMilli()
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(message)
	}
	end := time.Now().UnixMilli()

	fmt.Println(end - start)
}
