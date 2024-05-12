package main

import "fmt"

func main() {
	var car vehicle;
	car.Name = "Tesla";
	car.wheels = 4;
	car.drive();
}

type vehicle struct {
	Name string
	wheels int
}

func (a vehicle) drive() {
	fmt.Println("Driving");
}

