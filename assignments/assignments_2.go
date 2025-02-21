package assignments

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// STRUCTURES

// 1. ARRAY #1
// Create a program where the user is asked four questions to input a number.
// Save all the numbers in a list. Loop through the list and find the largest number.
// Print the result back to the screen for the user.

func FindGreatestNum() {
	numbers := [4]int{}
	for i := 0; i < 4; i++ {
		inputValue := intInput("Enter a number: ")
		numbers[i] = inputValue
	}
	fmt.Println("You entered: ", numbers)

	greatestNum := numbers[0]
	for _, num := range numbers {
		if num > greatestNum {
			greatestNum = num
		}
	}

	fmt.Printf("Out of all numbers %v was the greatest.", greatestNum)

}

// 2. SLICE #2
// Continuously ask for names to be input and store them in a slice.
// Ask if the user wants to continue (y or n).
// When the user chooses not to continue, print out the longest name (most characters).

// PERSONAL NOTE FOR IMPROVEMENT: add slice to store multiple names if two or more qualify as longest
func NameSlice() {
	var repeat string
	nameList := []string{}
	for {
		var name string
		fmt.Print("Enter a name to add it to the list: ")
		fmt.Scanln(&name)
		nameList = append(nameList, name)
		fmt.Println(nameList)
		fmt.Print("Do you want to continue (y or n): ")
		fmt.Scanln(&repeat)
		if repeat != "y" {
			break
		}
	}

	if len(nameList) == 0 {
		fmt.Println("No names entered.")
		return
	}

	longestName := nameList[0]
	for _, n := range nameList {
		if len(n) > len(longestName) {
			longestName = n
		}
	}
	fmt.Printf("Longest name in the list is %v with %v characters.", longestName, len(longestName))
}

// STRING HANDLING

// 1. STRING #1
// Ask the user to input 3 strings. Concatenate them into one long string (with spaces in between)
// and print the result to the screen.

func ConcatString() {
	stringValues := []string{}
	var input string
	for i := 0; i < 3; i++ {
		//var input string -> IMPROVED: Moved this outside the loop as it's unnecessary to recreate it each time
		fmt.Printf("Enter a word (%v)/3: ", i+1)
		fmt.Scan(&input)
		stringValues = append(stringValues, input)
	}
	joined := strings.Join(stringValues, " ")
	fmt.Println(joined)
}

// 2. STRING #2
// You have a string variable containing the following text, "Hello, world"
// - Find the first occurrence of the letter 'w' (which position in the string)
// - Find the last occurrence of the letter 'o' (which position in the string)

func LocateLetterPosition() {
	greeting := "Hello, world"
	position1 := strings.Index(greeting, "w")
	position2 := strings.LastIndex(greeting, "o")
	fmt.Printf("The first occurrence of 'w' is at index: %d, and the last occurrence of 'o' is at index: %d.", position1, position2)

}

// 3. STRING #3
// You have the string variable 'namn="kurt andersson"';
// Write code so that the first name and last name in the variable 'namn' are capitalized.
// The result should be: "Kurt Andersson"

func CapitalizeWords() {
	name := "kurt andersson"
	slice := strings.Fields(name)
	for i, word := range slice {
		x := strings.Split(word, "")
		x[0] = strings.ToUpper(x[0])
		joined := strings.Join(x, "")
		slice[i] = joined
	}
	result := strings.Join(slice, " ")
	fmt.Print(result)
}

// PERSONAL NOTE: USING RUNES SOLUTION
// While word[0] gives the byte representation of the character,
// wrapping it with rune() converts it into a rune, which represents the character in Unicode.

func CapitalizeWords2() {
	name := "kurt andersson"
	slice := strings.Fields(name)

	for i, word := range slice {
		slice[i] = string(unicode.ToUpper(rune(word[0]))) + word[1:]
	}

	result := strings.Join(slice, " ")

	fmt.Print(result)
}

// 4. STRING #4
// You have a string with the text "Detta är en sträng som du skall ändra".
// Replace all spaces in the string with the '*' character.
// Then calculate how many occurrences of '*' there are in the string.

func ReplaceContent() {
	phrase := "This is a string that needs to be changed."
	newPhrase := strings.ReplaceAll(phrase, " ", "*")
	characterRepeats := strings.Count(newPhrase, "*")
	fmt.Print(characterRepeats)
}

// 5. STRING #5
// Ask the user to input an email address.
// The program should check that the input is valid:
// that it contains an '@' symbol and that there is a '.' with 2 or 3 characters after it.
// Notify the user whether the address is correct or invalid.

func ValidateEmail() {
	var inputValue string
	fmt.Print("Enter your email: ")
	fmt.Scanln(&inputValue)

	if strings.Contains(inputValue, "@") && strings.Contains(inputValue, ".") {
		lastDotIndex := strings.LastIndex(inputValue, ".")
		suffix := inputValue[lastDotIndex+1:]
		if len(suffix) < 2 || len(suffix) > 3 {
			fmt.Printf("invalid entry '%v'", inputValue)
			return
		}
		fmt.Print("Email validated!")
	} else {
		fmt.Print("Email is missing either @ or .xxx")
	}
}

// 6. STRING #6
// Ask the user to input a word or sentence.
// The program should check if the word is a palindrome, meaning if the word is the same
// when read forwards and backwards.
// Example palindromes: "anna", "otto", or the sentence "ni talar bra latin".
// Spaces should be ignored. Ignore punctuation as well.

func CheckIfPalindrome() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a word or a sentence: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	var filteredRunes []rune

	for _, rune := range input {
		if unicode.IsLetter(rune) {
			filteredRunes = append(filteredRunes, unicode.ToLower(rune))
		}
	}

	length := len(filteredRunes)

	for i := range length / 2 {
		if filteredRunes[i] != filteredRunes[length-1-i] {
			fmt.Printf("\"%v\" is not a palindrome.", input)
			return
		}
	}
	fmt.Printf("\"%v\" is a palindrome.", input)

}

// PERSONAL NOTES:
// 1. OG FOR LOOP:
// for i := 0; i < length/2; i++ {
// 	if runes[i] != runes[length -1 -i] {
// 		fmt.Printf("%v is not a palindrome.", input)
// 		return
// 	}
// }
// update: replacing a 3-clause for i := 0; i < n; i++ {} loop by for i := range n {}, added in go1.22;

// 2.
// bufio.NewReader(os.Stdin) allows me to store complete sentences

// # ------------------------------------------------------ FUNCTIONS

// 1. Add
// Create a program with a new function.
// The method should take two parameters of type string and concatenate them into one string,
// then return the new value. Call the new method and print the result to the screen.

func Joined(a string, b string) string {
	word := a + " " + b
	return word
}

// 2. VAT (MOMS)
// Create a function that calculates the VAT (Value Added Tax) for a given amount.
// The amount should be a parameter of type float64.
// The method should return the VAT value.

func AddVAT(a float64) float64 {
	vat := a * 0.25
	sum := a + vat
	return sum
}

// 3. LONGEST WORD
// Create a method called HittaLangstaOrdet.
// The method should take a slice of strings as a parameter.
// It should loop through the array and return the longest word.
// Is there a way to do this without needing to create a local longest variable or even using return?
// Answer: no, it's not possible since I have to compare each element to something. also if the function doesn't return anything, I wont be able to store it inside a variable. I could still call on the function but it wouldn't do anything.

func FindLongestWord(a []string) string {
	var longestWord string
	if len(a) != 0 {
		longestWord = a[0]
		for _, word := range a {
			if len(word) > len(longestWord) {
				longestWord = word
			}
		}
	}
	return longestWord
}

// 4. CALCULATE TAX
// Create a function called CalculateTaxesOnSalary
// The method takes a salary (number) as an input parameter.
// If the monthly salary >= 30,000, the tax is salary * 0.33
// If the monthly salary < 15,000, the tax is salary * 0.12
// If the monthly salary is between 15,000 and 30,000, the tax is salary * 0.28
// The method should return the calculated tax amount.
// It should also return an error code (or nil) if the salary is negative (invalid input).

// OG thoughts:
// func CalculateTaxesOnSalary(salary float64) float64 || nil {
// 	var vat float64
// 	if salary >= 30000 {
// 		vat = salary * 0.33
// 	} else if salary < 15000 {
// 		vat = salary * 0.12
// 	} else if salary >= 15000 && salary = 30000 {
// 		vat = salary * 0.28
// 	} else {
// 		return nil
// 	}
// 	return salary + vat
// }

func CalculateTaxesOnSalary(salary float64) *float64 {
	var vat float64
	if salary >= 30000 {
		vat = salary * 0.33
	} else if salary < 15000 {
		vat = salary * 0.12
	} else if salary >= 15000 && salary < 30000 {
		vat = salary * 0.28
	} else {
		return nil
	}
	total := salary + vat
	return &total
}
