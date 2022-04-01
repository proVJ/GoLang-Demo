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

	// initializeVariable()

	// howToCallpackageLevelVariable()

	// howToCallGlobalLevelVariable()

	// howToCallOtherPackageMethod()

	// constantImplementation()

	// arrayDeclaration()

	mapDeclarationAndImplementation()

}

//KEY VALUE PAIR
func mapDeclarationAndImplementation() {
	fmt.Println("-----------Key Value Pair Implementation-----------")
	// var emp = make(map[string]int) // declaration
	var emp = map[string]int{
		"Emp 1": 1,
		"Emp 2": 2,
	}
	fmt.Println(emp)

	fmt.Println("Search by Key ")
	fmt.Println(emp["Emp 1"])

	fmt.Println("Update value of Emp 1")
	emp["Emp 1"] = 1000
	fmt.Println(emp["Emp 1"])

	fmt.Println("Add a new Record")
	emp["Emp 3"] = 4000
	fmt.Println(emp)
}

// ARRAY
func arrayDeclaration() {
	fmt.Println("-----------Array Implementation-----------")
	var array [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(array) //  array[0]
	fmt.Println(len(array))

	fmt.Println("-----------N Number Array Implementation-----------")
	var nArray = [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nArray) //  array[0]
	fmt.Println(len(nArray))

}

//Constant
func constantImplementation() {
	fmt.Println("-----------Constant Implementation-----------")
	const d int = 45
	// fmt.Printf("Constant Value : %v", d)
	fmt.Println(d)

}

//Calling Hello method from greetings folder
func howToCallGlobalLevelVariable() {
	fmt.Println("-----------CALLING Global LEVE VARIABLE-----------")
	fmt.Println(greetings.GlobalVariable)
}

//Calling Hello method from greetings folder
func howToCallOtherPackageMethod() {
	fmt.Println("-----------CALLING METHOD FROM DIFFERENT FILE-----------")
	greetings.Hello("greetings.go greetings method")

	//Calling PrintHello method from hello folder
	hello.PrintHello()
}

//GLOBAL VARIABLE
func howToCallpackageLevelVariable() {
	fmt.Println("-----------GLOBAL VARIABLE-----------")
	fmt.Println(packageLevelVariable)
}

// 	LOCAL VARIABLE
func initializeVariable() {
	fmt.Println("-----------INITIALIZE-----------")

	var intVariable int // Declaration
	intVariable = 45245 // Initialization
	fmt.Println(intVariable)

	var myVariable = "Hello, Variables" // Declaration with Initialization
	fmt.Println(myVariable)
}
