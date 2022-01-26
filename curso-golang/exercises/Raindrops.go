package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Enter the raindrops number: ")
	var number int
	fmt.Scanf("%s", number)
	var response string
	if i := number % 3; i == 0 {
		fmt.Println(i)
		response += "Pling"
	} else if i := number % 5; i == 0 {
		response += "Plang"
	} else if i := number % 7; i == 0 {
		response += "Plong"
	} else {
		response = strconv.Itoa(number) + " is not factored by 3, 5, or 7, so the result would be" + strconv.Itoa(number)
	}

	fmt.Println(response)
}
