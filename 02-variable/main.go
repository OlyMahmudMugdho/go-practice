package main;

import "fmt";

func main () {
	var name string = "Mugdho";
	fmt.Printf("The type of name is %T \n", name);

	var age int = 23;
	fmt.Printf("The type of age is %T \n", age);

	// implicit type
	salary := 50000;
	fmt.Println("My salary is ", salary);

	const pi float32 = 3.14; // Constant
	fmt.Println("The value of Pi is ", pi);
	
}