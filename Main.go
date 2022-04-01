package main

import (
	"fmt"

	"greetings"
	"hello"
)

func main() {
	// Before running this Please Read Me File

	var intVariable int // Declaration
	intVariable = 45245 // Initialization
	fmt.Println(intVariable)

	var myVariable = "Hello, Variables" // Declaration wit Initialization
	fmt.Println(myVariable)

	//Calling Hello method from greetings folder
	greetings.Hello("greetings.go greetings method")

	//Calling PrintHello method from hello folder
	hello.PrintHello()
}
