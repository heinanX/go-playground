package main

import (
	"fmt"
	//"heinanx20250204/assignments"
	"heinanx20250204/basics"
)

func main() {

	// ---------------------------------------------> ASSIGNMENTS 3

	//assignments.Structs()
	basics.IntroductionAboutStructs()
	dog3 := basics.NewDog("kodybarks", "baloney", "finnish lapphound")
	fmt.Println(dog3)

	// ---------------------------------------------> ASSIGNMENTS 2
	//assignments.CheckIfPalindrome()

	// words := []string{"dog", "giraffe", "dolphin", "rat", "crocodile"}
	// printThis := assignments.FindLongestWord(words)

	// printThis := assignments.CalculateTaxesOnSalary(25000)
	// if printThis != nil {
	// 	fmt.Println("Total Salary with tax:", *printThis) // The asterisk dereferences the pointer to get the value that the pointer is pointing to, without it, it would only print the memory address
	// } else {
	// 	fmt.Println("Invalid salary")
	// }

	// fmt.Print(printThis)

}
