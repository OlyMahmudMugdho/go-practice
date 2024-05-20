package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5};

	fmt.Println(nums)
	
	// removing index 2
	nums = append(nums[:2], nums[2+1:]... )

	fmt.Println(nums)
}