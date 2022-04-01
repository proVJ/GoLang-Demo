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

	constantImplementation()

}

func constantImplementation() {
	fmt.Println("-----------Constant Implementation-----------")
	const d int = 45
	fmt.Printf("Constant Value : %v", d)
}

//==========CALLING PACKAGE LEVE VARIABLE ==================
//Calling Hello method from greetings folder
func howToCallGlobalLevelVariable() {
	fmt.Println("-----------CALLING Global LEVE VARIABLE-----------")
	fmt.Println(greetings.GlobalVariable)
}

//==========CALLING METHOD FROM DIFFERENT FILE ==================
//Calling Hello method from greetings folder
func howToCallOtherPackageMethod() {
	fmt.Println("-----------CALLING METHOD FROM DIFFERENT FILE-----------")
	greetings.Hello("greetings.go greetings method")

	//Calling PrintHello method from hello folder
	hello.PrintHello()
}

//==========CALL GLOBAL VARIABLE ==================
func howToCallpackageLevelVariable() {
	fmt.Println("-----------GLOBAL VARIABLE-----------")
	fmt.Println(packageLevelVariable)
}

//==========INITIALIZE ==================
// 	LOCAL VARIABLE
func initializeVariable() {
	fmt.Println("-----------INITIALIZE-----------")

	var intVariable int // Declaration
	intVariable = 45245 // Initialization
	fmt.Println(intVariable)

	var myVariable = "Hello, Variables" // Declaration with Initialization
	fmt.Println(myVariable)
}
