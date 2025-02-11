package main

import "fmt"

func main() {
	//possible way to declare a variable but not used
	// var age = 34

	// possible way to declare a variable and assign a value later
	//var age int
	//... code
	//age = 34

	age := 34 //declares and sets value at the same timeL

	name := "Linda"

	fmt.Printf("Hello %s %d years old", name, age)
}
