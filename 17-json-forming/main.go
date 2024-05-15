package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	taskOne := Todo{1, 1, "something", false}
	todoJson, _ := json.MarshalIndent(taskOne, "", "\t")
	fmt.Println(string(todoJson))

}
