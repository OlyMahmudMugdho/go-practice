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

type User struct {
	Id       int
	Name     string
	Roll     int
	password string
	// password will not included in json
}

func main() {

	taskOne := Todo{1, 1, "something", false}
	todoJson, _ := json.MarshalIndent(taskOne, "", "\t")
	fmt.Println(string(todoJson))

	user1 := User{1, "Mugdho", 14, "abc"}
	userJson, _ := json.Marshal(user1)
	fmt.Println(string(userJson))

}
