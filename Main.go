package main

import (
	"fmt"

	"greetings"
	"hello"
)

// GLOBAL VARIABLE
var globalVariable int = 8784524

func main() {
	// BEFORE RUNNING THIS PLEASE READ README.MD FILE

	//==========HOW TO INITIALIZE ==================
	// 	LOCAL VARIABLE

	var intVariable int // Declaration
	intVariable = 45245 // Initialization
	fmt.Println(intVariable)
	fmt.Println("====={HOW TO INITIALIZE}=====")
	var myVariable = "Hello, Variables" // Declaration wit Initialization
	fmt.Println(myVariable)

	//==========HOW TO CALL GLOBAL VARIABLE ==================
	fmt.Println("====={GLOBAL VARIABLE}=====")
	fmt.Println(globalVariable)

	//==========CALLING METHOD FROM DIFFERENT FILE ==================
	//Calling Hello method from greetings folder
	fmt.Println("====={CALLING METHOD FROM DIFFERENT FILE}=====")
	greetings.Hello("greetings.go greetings method")

	//Calling PrintHello method from hello folder
	hello.PrintHello()
}
