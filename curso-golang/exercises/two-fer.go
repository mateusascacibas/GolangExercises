package main

import "fmt"

//Two-fer or 2-fer is short for two for one. One for you and one for me. Where "name" is the given name.
// However, if the name is missing, return the string

func main() {
	fmt.Println("Enter Your Name: ")

	var name string
	var response string
	fmt.Scanf("%s", &name)
	if name != "" {
		response = "One for " + name + ", one for me."
	} else {
		response = "One for you, one for me."
	}

	fmt.Println(response)
}
