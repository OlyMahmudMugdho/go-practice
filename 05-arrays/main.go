package main

import "fmt"


func main()  {
	var nums [4]int;
	nums[0] = 5;
	nums[1] = 4;
	nums[2] = 8;
	nums[3] = 9;
	
	fmt.Println(nums)
	fmt.Println("the length of nums array is ", len(nums));

	var names = [3]string{"Mila","Mugdho","Meghla"};
	fmt.Println(names);
	
	var langs [3]string;

	langs[0] = "go"
	langs[1] = "java"
	
	fmt.Println(langs)

}