package assignments

import (
	"fmt"
	"strings"
)

// 1. Write a console application that prints Hello World
func PrintGreeting() {
	fmt.Println("Hello World")
}

// 2. Expressions and Data types
// Write a, application that does the following:
// 	print(“2021”)
//	print(2021)
//	print(“2021-12-24”)
//	print(2021-12-24)
// What are the results? And why?

func PrintDate() {
	print("2021")         // this is a string
	print(2021)           // this is a number (int)
	print("2021-12-24")   // this is a string
	print(2021 - 12 - 24) // this is simple math, not a date. It subtracts 12 and 24 from 2021, resulting in 1985
}

//Explainantion: Everything is written in one long sentence since there are no line breaks. This would be avoidable by using Println function from the fmt package.

// VARIABLES

// 3. String and int
// "Create a string variable name with your name.
//  Create an int variable age with your age.
//  Then print 'My name is Kalle (the content of name) and I am 27 (the content of age) years old.'"

func StringAndInt() {
	name := "Linda"
	age := 34
	fmt.Printf("My name is %s and I'm %d years old", name, age)
}

// 4. First Name and Last Name
// Create an application where the user enters their first and last name.

//  Print "Enter your first name:". Capture the value in a variable.
//  Have the cursor wait for your input on the same line.
//  Do the same for the last name.
//  Then print the names in reverse order.
//  Make sure the result looks like this:
// 	  Enter your first name: Jane
//	  Enter your last name: Doe
//	  Your name is: Doe, Jane

func FirstAndLastName() {
	var firstname string
	var lastname string

	fmt.Print("Enter your first name: ")
	fmt.Scanln(&firstname)

	fmt.Print("Enter your last name: ")
	fmt.Scanln(&lastname)

	fmt.Printf("Your name is: %s, %s", lastname, firstname)
}

// 5. String Operators

// Ask the user to input their name. Store it in a variable called name.
// Create a new variable nameRepeated with the name repeated five times in a row. Print the new variable.
// Example output:
//	 What is your name: Stefan
//	 StefanStefanStefanStefanStefan

func PrintNameFiveTimes() {
	var name string

	fmt.Print("What's your name: ")
	fmt.Scanln(&name)

	nameRepeated := strings.Repeat(name, 5)

	fmt.Print(nameRepeated)
}

// IF-STATEMENTS

// Ask the user to enter a number.
// Check if it's greater than 10.
// If it is, print "The number is greater than 10."
// If it's not, print "The number is less than 10."
