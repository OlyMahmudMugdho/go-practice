package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Enter your age : ");

	reader := bufio.NewReader(os.Stdin);

	ageInput,_ := reader.ReadString('\n');

	age,_ := strconv.ParseInt(strings.TrimSpace(ageInput),10, 64);

	fmt.Println("in sha Allah, next year your age will be ", age + 1);


}