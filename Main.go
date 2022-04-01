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

//==========CALLING PACKAGE LEVE VARIABLE ==================
//Calling Hello method from greetings folder
func howToCallGlobalLevelVariable() {
	fmt.Println("====={CALLING Global LEVE VARIABLE}=====")
	fmt.Println(greetings.GlobalVariable)
}

//==========CALLING METHOD FROM DIFFERENT FILE ==================
//Calling Hello method from greetings folder
func howToCallOtherPackageMethod() {
	fmt.Println("====={CALLING METHOD FROM DIFFERENT FILE}=====")
	greetings.Hello("greetings.go greetings method")

	//Calling PrintHello method from hello folder
	hello.PrintHello()
}

//==========HOW TO CALL GLOBAL VARIABLE ==================
func howToCallpackageLevelVariable() {
	fmt.Println("====={GLOBAL VARIABLE}=====")
	fmt.Println(packageLevelVariable)
}

//==========HOW TO INITIALIZE ==================
// 	LOCAL VARIABLE
func initializeVariable() {
	var intVariable int // Declaration
	intVariable = 45245 // Initialization
	fmt.Println(intVariable)
	fmt.Println("====={HOW TO INITIALIZE}=====")
	var myVariable = "Hello, Variables" // Declaration wit Initialization
	fmt.Println(myVariable)
}
