package main

import "fmt"

func main() {
	languages := make(map[string]string); // map[key_type]value_type

	languages["go"] = "Golang"
	languages["js"] = "JavaScript"
	languages["cs"] = "C#"
	languages["java"] = "Java"

	fmt.Println(languages);

	frameworks := make(map[int]string);

	frameworks[0] = "Spring Boot";
	frameworks[1] = "ASP.NET Core";
	frameworks[2] = "Express.js";
	
	fmt.Println(frameworks)

	// deleting

	delete(languages,"js");
	fmt.Println("After deleting JS")
	fmt.Println(languages)
}