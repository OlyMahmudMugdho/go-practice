package main

import "fmt"

func main() {
	var langs = []string{}; // initialized a slice
	fmt.Println(langs)

	// adding
	langs = append(langs, "go");
	fmt.Println(langs)
	langs = append(langs, "java", "c#");
	fmt.Println(langs);

	// removal

	langs = append(langs[0:2])  // removing c#
	fmt.Println(langs)
	

}