package assignments

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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

// # ---------------------------------------------------	VARIABLES

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

// # ---------------------------------------------------	IF-STATEMENTS

// 1. Greate than ten
// Ask the user to enter a number.
// Check if it's greater than 10.
// If it is, print "The number is greater than 10"
// If it's not, print "The number is less than 10"

func GreaterThanTen() {
	var num int
	var text string

	fmt.Print("Enter a Number: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("It doesn't seem to be a number.")
		return
	}

	if num > 10 {
		text = "The number is greater than 10"
	} else {
		text = "The number is less than 10"
	}
	fmt.Println(text)

}

// VALIDATION FUNCTION FOR FUTURE INT INPUTS BECASUE I'M GETTING TIRED OF WRITING THE SAME LOGIC
func intInput(text string) int {
	var input int
	for {
		fmt.Print(text)
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("It doesn't seem to be a number. Please try again.")
			fmt.Scanln() // Clears the input buffer
		} else {
			return input
		}
	}
}

// 2. Order Milk
// Ask the user how many cartons of milk are left.
// If it's less than 10, print "Order 40 more."
// If it's between 10 and 20, print "Order 20 more."
// Otherwise, print "You don't need to order any milk."

func OrderMilk() {
	num := intInput("How many cartons of milk are left? ")

	if num < 10 {
		fmt.Println("Order 40 more.")
	} else if num >= 10 && num <= 20 {
		fmt.Println("Order 20 more.")
	} else {
		fmt.Println("You don't need to order any milk.")
	}

}

// 3. Largest Number
// Ask the user to enter two numbers: num1 and num2.
// Store the largest number of the two in a third variable called largest.
// Print: The largest number is <>

func LargestNum() {
	var largest int

	num1 := intInput("Enter a number: ")

	num2 := intInput("Enter another number: ")

	if num1 > num2 {
		largest = num1
	} else {
		largest = num2
	}

	fmt.Printf("The largest number is %d", largest)
}

// 4. Largest Number
// repeat task 3 but with three inputs.
func LargestNum2() {
	var largest int

	num1 := intInput("Enter a number: ")

	num2 := intInput("Enter a second number: ")

	num3 := intInput("Enter a third number: ")

	if num1 >= num2 && num1 >= num3 { //Improvement: By checking if they are greater or _equal_ to the other two numbers, we eliminate the risk of running unnecessary code.
		largest = num1
	} else if num2 >= num1 && num2 >= num3 {
		largest = num2
	} else {
		largest = num3
	}

	fmt.Printf("The largest number is %d", largest)
}

// 5. Fare Calculator
// Ask the user to enter which category they belong to: adult, senior, or student.
//   If they are a senior or a student, the trip costs 20 SEK.
//   If they are an adult, the trip costs 30 SEK.
//   Otherwise, inform the user that they have entered an incorrect category.

// Original Thought proscess:
// func CalcFare() {
// 	var cat string

// 	fmt.Print("Enter your category (student, senior, adult): ")
// 	fmt.Scanln(&cat)

// 	if cat == "senior" || cat == "student" {
// 		fmt.Println("Your fare is 20 SEK")
// 	} else if cat == "adult" {
// 		fmt.Println("Your fare is 30 SEK")
// 	} else {
// 		fmt.Println("Your input is invalid")
// 		CalcFare()
// 	}
// }

// Disadvantage: Every time a function calls itself, there’s overhead involved—allocating memory, maintaining the stack, etc.
// It may also cause stack overflow as the task is added to the stack everytime an invalid response is given.
// For simple tasks like user input validation, loops are more efficient and readable.

// Improved:
func CalcFare() {
	var cat string

	for {
		fmt.Print("Enter your category (student, senior, adult): ")
		fmt.Scanln(&cat)

		if cat == "senior" || cat == "student" {
			fmt.Println("Your fare is 20 SEK")
			break
		} else if cat == "adult" {
			fmt.Println("Your fare is 30 SEK")
			break
		} else {
			fmt.Println("Your input is invalid. Please enter 'student', 'senior', or 'adult'.")
		}
	}
}

// 6. Enter your birth year
//	 Ask the user to enter their birth year.
//   If it is greater than or equal to 1980 but less than 1990, print:
//   "You were born in the 1980s."
//   If it is greater than or equal to 1990 but less than 2000, print:
//   "You were born in the 1990s."
//   If it is less than 1980 or greater than or equal to 2000, print:
//   "You were not born in the 1980s or 1990s."

func CheckBirthYear() {
	birthyear := intInput("Enter your year of birth? ")

	if birthyear >= 1980 && birthyear < 1990 {
		fmt.Println("You were born in the 1980s!")
	} else if birthyear >= 1990 && birthyear < 2000 {
		fmt.Println("You were born in the 1990s!")
	} else {
		fmt.Println("You were not born in the 1980s or 1990s.")
	}
}

// 7. Are you Scandinavian?
// 	Ask the user for their country.
//  If it's Sweden, Denmark, or Norway, print: "You live in Scandinavia."
//  Otherwise, print: "You don't live in Scandinavia."

func IsScandinavian() {
	var nationality string

	fmt.Print("Where do you live? ")
	fmt.Scanln(&nationality)

	nationality = strings.ToLower(nationality)

	if nationality == "sweden" || nationality == "denmark" || nationality == "norway" {
		fmt.Println("You live in Scandinavia.")
	} else {
		fmt.Println("You don't live in Scandinavia.")
	}
}

// 8. Fast Food Order
//  Ask the user how much money they have. Then, ask if they have a discount.
//  If the amount is between 15 and 25 and the user does not have a discount, print: "You can buy a small hamburger."
//  If the amount is between 15 and 25 and the user has a discount, print: "You can buy a small hamburger and fries."
//  If the amount is greater than 25 and less than or equal to 50 and the user does not have a discount, print: "You can buy a large hamburger."
//  If the amount is greater than 25 and less than or equal to 50 and the user has a discount, print: "You can buy a large hamburger and fries."
//  If the amount is greater than 60 or if it is between 50 and 60 and the user has a discount, print: "You can buy a meal with a drink."

// func FastFoodOrder() {
// 	var discount string
// 	money := intInput("How much money do you have? ")
// 	fmt.Print("Are you eligible for a discount? (yes/no)")
// 	fmt.Scanln(&discount)
// 	if discount != "yes" && money >= 15 && money <= 25 {
// 		fmt.Println("You can buy a small hamburger.")
// 	} else if discount == "yes" && money >= 15 && money <= 25 {
// 		fmt.Println("You can buy a small hamburger and fries.")
// 	} else if discount != "yes" && money > 25 && money <= 50 {
// 		fmt.Println("You can buy a large hamburger.")
// 	} else if discount == "yes" && money > 25 && money <= 50 {
// 		fmt.Println("You can buy a large hamburger and fries.")
// 	} else if discount == "yes" && money > 50 {
// 		fmt.Println("You can buy a meal with a drink.")
// 	} else {
// 		fmt.Println("Sorry, bro. You need to save up some more!")
// 	}
// }

// PERSONAL NOTES ON IMPROVEMENT:

func FastFoodOrder() {
	var discount string
	money := intInput("How much money do you have? ")
	fmt.Print("Are you eligible for a discount? (yes/no) ")
	discount = strings.ToLower(discount)
	fmt.Scanln(&discount)

	if money >= 15 && money <= 25 {
		if discount == "yes" { // if statement inside the if statement prevents rewriting same code
			fmt.Println("You can buy a small hamburger and fries.")
		} else {
			fmt.Println("You can buy a small hamburger.")
		}
	} else if money > 25 && money <= 50 {
		if discount == "yes" {
			fmt.Println("You can buy a large hamburger and fries.")
		} else {
			fmt.Println("You can buy a large hamburger.")
		}
	} else if money > 60 || (money >= 50 && discount == "yes") {
		fmt.Println("You can buy a meal with a drink.")
	} else {
		fmt.Println("Sorry, bro. You need to save up some more!")
	}
}

// # ---------------------------------------------------	LOOPITY LOOPS

// 1. LOOP #1
// Create a program that prints the numbers from 0 to 10 on the screen.

func LoopyTens() {
	n := 0
	for n <= 10 {
		fmt.Print(n)
		n++
	}
}

// 2. LOOP #2
// Create a program where the user inputs two numbers.
// Then, have the program print all the numbers between these two numbers on the screen.

func NumberInbetween() {
	num1 := intInput("Enter start number: ")
	num2 := intInput("Enter end number: ")
	var start int
	var end int

	if num1 < num2 {
		start = num1
		end = num2
	} else {
		start = num2
		end = num1
	}
	i := start
	for i < end {
		fmt.Printf("%d ", i)
		i++
	}
}

// 3. LOOP #3
// Create a program where the user:
// - Inputs two numbers.
// - Then, print the sum of the two numbers.
// - Then, ask the question: "Do you want to continue? (Y/N)".
// - If the user answers "Y", repeat steps a-c.
// - If the user answers "N", stop the program.

func AddNumbers() {
	for {
		num1 := intInput("Enter number: ")
		num2 := intInput("Enter second number: ")
		var repeat string

		fmt.Println("Sum: ", num1+num2)
		fmt.Print("Do you want to continue? (Y/N)")
		fmt.Scan(&repeat)
		repeat = strings.ToLower(repeat)
		if repeat != "y" {
			break
		}
	}
}

// 4. LOOP #4
// Ask the user to input a number. Save the value in a variable.
// Repeat this 10 times. For each input, add the entered value to the variable's value.
// When done, print: "The sum of what you entered is: <sum>."

func Number4() {
	sum := 0
	for i := 1; i <= 10; i++ {
		userInput := intInput("Enter a number: ")
		sum = sum + userInput // (improvement: sum += userInput)
	}
	fmt.Println(sum)

}

// 5. LOOP #5
// Create a program where the user inputs a number.
// Then, have the program print all the numbers less than the entered number but greater than 0.
// Solve this using a loop.

func PrintRange() {

	userInput := intInput("Enter a number: ")
	for i := userInput - 1; i > 0; i-- {
		fmt.Println(userInput)
	}

}

// 6. Rolling dice
//Roll two dice and display the result. Repeat until the user does not answer 'y' or 'yes'.
// PERSONAL THOUGHTS ON IMPROVEMENT: Added time delay between prints

func DiceRoll() {
	var repeat string
	for {
		fmt.Println("Rolling the dice...")
		time.Sleep(1 * time.Second)
		fmt.Println("The values are...")
		time.Sleep(1 * time.Second)
		fmt.Println(rand.Intn(6) + 1)
		fmt.Println(rand.Intn(6) + 1)
		fmt.Print("Roll the dice again? Y/N ")
		fmt.Scan(&repeat)

		repeat = strings.ToLower(repeat)

		if repeat != "y" {
			break
		}
	}
}
