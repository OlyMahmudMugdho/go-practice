package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var results = []map[int]int{}

func main() {
	var nums = []int{1, 2, 3, 4, 5}
	for _, v := range nums {
		go getFactorials(v)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(results)
}

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

func getFactorials(num int) {
	defer wg.Done()
	mut.Lock()
	res := factorial(num)
	results = append(results, map[int]int{num: res})
	mut.Unlock()
}
