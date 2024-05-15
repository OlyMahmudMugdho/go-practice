package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	UserId    int `json:"userId`
	Id        int	
	Title     string
	Completed bool
}


func main() {

	taskOne  := Todo{1,1,"something", false};
	todoJson,_ := json.Marshal(taskOne);
	fmt.Println(string(todoJson))
	
}
