package basics

import (
	"fmt"
)

func IfStatements() {
	var name string
	var age string

	fmt.Print("What's your name?")
	fmt.Scanln(&name) //the & works as a pointer to the object we want to assign the value to. It contains the "address"

	fmt.Print("What's your age?")
	_, err := fmt.Scanln(&age)
	if err != nil {
		fmt.Println("It doesn't seem to be a number.")
		return
	}
}
