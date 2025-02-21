package basics

import (
	"fmt"
)

// Structs are defined at the package level
type Dog struct {
	name     string
	nickname string
	breed    string
	age      int8
}

// function is not tied to the Dog struct itself
func printNickname(d Dog) {
	fmt.Println(d.nickname)
}

// a way to make a "traditional" method
func (d Dog) printName() {
	fmt.Println(d.name)
}

// a kind of Constructor ---> common practice
func NewDog(dname string, dnickname string, dbreed string) Dog {
	d := Dog{name: dname, nickname: dnickname, breed: dbreed}
	return d
}

func IntroductionAboutStructs() {
	dog := Dog{name: "Buddy", nickname: "Bud", breed: "Golden Retriever", age: 3}
	dog2 := NewDog("koda", "kody", "finnish lapphound")
	dog.printName()
	printNickname(dog)
	dog2.printName()
}
