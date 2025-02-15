package basics

import (
	"fmt"
	"strings"
)

func PrintString() {
	//possible way to declare a variable but not used
	// var age = 34

	// possible way to declare a variable and assign a value later
	//var age uint8 (uint8 for better performance as it takes less memory than default int)
	//... code
	//age = 34

	age := 34 //declares and sets value at the same timeL
	name := "Carly"

	// specifying the value %s = string, and %d = digit
	fmt.Printf("Hello %s. You seem to be %d years old\n", name, age)
	// to simplify we may instead use %v (=value)
	fmt.Printf("Hello %v. You seem to be %v years old\n", name, age)

	if strings.Contains(strings.ToLower(name), "carly") {
		fmt.Println("That's a good name")
	} else {
		fmt.Println("You should change")
	}

}
