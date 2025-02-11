package main

import (
	"fmt"
	"strings"
)

func main() {
	//possible way to declare a variable but not used
	// var age = 34

	// possible way to declare a variable and assign a value later
	//var age uint8 (uint8 for better performance as it takes less memory than default int)
	//... code
	//age = 34

	age := 34 //declares and sets value at the same timeL
	name := "Carly"

	fmt.Printf("Hello %s. You seem to be %d years old\n", name, age)

	if strings.Contains(name, "Carly") {
		fmt.Println("That's a good name")
	} else {
		fmt.Println("You should change")
	}
}
