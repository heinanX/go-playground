package basics

import "fmt"

func LoopingLouie() {
	//OG FOR LOOP
	sum := 0

	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	// WHILE LOOP - while a < 10, write a and add 1
	a := 5
	for a < 10 {
		fmt.Println(a)
		a++
	}
}
