package main

import (
	"fmt"

	"greetings"
	"hello"
)

// PACKAGE LEVE VARIABLE (calmel case)
// this variable will accessible to this package only
var packageLevelVariable int = 8784524

func main() {
	// BEFORE RUNNING THIS PLEASE READ README.MD FILE

	initializeVariable()
	howToCallpackageLevelVariable()
	howToCallGlobalLevelVariable()
	howToCallOtherPackageMethod()

}

func howToCallGlobalLevelVariable() {
	//==========CALLING PACKAGE LEVE VARIABLE ==================
	//Calling Hello method from greetings folder
	fmt.Println("====={CALLING Global LEVE VARIABLE}=====")
	fmt.Println(greetings.GlobalVariable)
}

func howToCallOtherPackageMethod() {
	//==========CALLING METHOD FROM DIFFERENT FILE ==================
	//Calling Hello method from greetings folder
	fmt.Println("====={CALLING METHOD FROM DIFFERENT FILE}=====")
	greetings.Hello("greetings.go greetings method")

	//Calling PrintHello method from hello folder
	hello.PrintHello()
}

func howToCallpackageLevelVariable() {
	//==========HOW TO CALL GLOBAL VARIABLE ==================
	fmt.Println("====={GLOBAL VARIABLE}=====")
	fmt.Println(packageLevelVariable)
}

func initializeVariable() {
	//==========HOW TO INITIALIZE ==================
	// 	LOCAL VARIABLE
	var intVariable int // Declaration
	intVariable = 45245 // Initialization
	fmt.Println(intVariable)
	fmt.Println("====={HOW TO INITIALIZE}=====")
	var myVariable = "Hello, Variables" // Declaration wit Initialization
	fmt.Println(myVariable)
}
