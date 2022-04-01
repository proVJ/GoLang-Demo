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

	// mapDeclarationAndImplementation()

	// slicingArray()

	// sturctImplementation()

	// loopImplementation()

	// switchImplementation()

	interfaceImplementation()

	// deferImplementation()

	// functions(1, 2)

	//retuns a value
	// sum := nNumberParamerterSum(1, 2, 3)
	// fmt.Println("SUM", sum)

	//retuns multiple value
	// sum, sub := multipleReturnValue(10, 8)
	// fmt.Println("SUM ", sum, "Sub", sub)

	//Anonymous Function
	// f := func() {
	// 	fmt.Println("Anonymous Function")
	// }
	// f()

	// InterFace Implementation
	var obj StudentInterface = StudentStruct{
		Name: "Student 1",
		Id:   1,
	}

	obj.ShowDetails()
	fmt.Println(obj.ShowName())

}

func multipleReturnValue(i1, i2 int) (int, int) {
	sum := i1 + i2
	sub := i1 - i2
	return sum, sub

}

func nNumberParamerterSum(vals ...int) int {
	var sum int = 0
	for i := 0; i < len(vals); i++ {
		sum = sum + vals[i]
	}

	// for _, n := range vals {
	// 	sum += n
	// }
	return sum

}

func functions(x int, y int) {
	fmt.Println("SUM", (x + y))
}

//Defer
func deferImplementation() {
	fmt.Println("-----------Defer-----------")

	defer fmt.Println("------------------------------------")
	defer fmt.Println("-----------Close Connection-----------")

	fmt.Println("Connection Open")
	fmt.Println("Execute data")
	fmt.Println("Execute Transaction")

}

//=======================Inter Face anb Struct==================
// its defines the behavior
type StudentInterface interface {
	ShowDetails()
	ShowName() string
}

// its define the fields and properdy
type StudentStruct struct {
	Name string
	Id   int
}

func (stud StudentStruct) ShowDetails() {
	fmt.Println("Student Name", stud.Name, " Student ID", stud.Id)
}

func (stud StudentStruct) ShowName() string {
	return stud.Name
}

// Interface
func interfaceImplementation() {
	fmt.Println("-----------Interface-----------")

	var x interface{} = "test"
	fmt.Println(x)
	fmt.Println("------------------------------------")

}

// SWITCH
func switchImplementation() {
	fmt.Println("-----------SWITCH-----------")
	// data := 3
	switch data := 3; data {
	case 1:
		fmt.Println(data)
	case 2, 3:
		// case 2:
		// case 3:
		{
			fmt.Println("data can be 2 or 3, the value is", data)
		}
	default:
		fmt.Println("None, the value is ", data)
	}
	fmt.Println("------------------------------------")

}

// LOOP
func loopImplementation() {
	fmt.Println("-----------LOOP-----------")

	for i := 0; i < 10; i++ {
		fmt.Printf("Loop Value - %v \n", i+1)
	}
	fmt.Println("------------------------------------")

}

// Implementing Struct
type Employee struct {
	EmpId      int
	EmpName    string
	EmpMobiles []string
}

func sturctImplementation() {
	fmt.Println("-----------Stuct or Class-----------")
	var emp Employee = Employee{
		EmpId:      1,
		EmpName:    "Vijay",
		EmpMobiles: []string{"8784574541", "7874574215"},
	}
	fmt.Println(emp)
	fmt.Println("------------------------------------")

}

//Array Slicing
func slicingArray() {
	fmt.Println("-----------Previous Initialization-----------")
	// array
	// var arr [4] int = [4] int{1,2,3,4}
	var arr = [...]int{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Println("------------------------------------")

	fmt.Println("-----------Array Slicint-----------")
	//Slicing
	var arr2 []int = []int{1, 2, 3, 4}
	fmt.Println(arr2)
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

	fmt.Println("------------------------------------")
	fmt.Println("Search by Key ")
	fmt.Println(emp["Emp 1"])
	fmt.Println("------------------------------------")

	fmt.Println("Update value of Emp 1")
	emp["Emp 1"] = 1000
	fmt.Println(emp["Emp 1"])
	fmt.Println("------------------------------------")

	fmt.Println("Add a new Record")
	emp["Emp 3"] = 4000
	fmt.Println(emp)
	fmt.Println("------------------------------------")

	fmt.Println("Delete one Record")
	delete(emp, "Emp 3")
	fmt.Println(emp)
	fmt.Println("------------------------------------")

	fmt.Println("Check Existence Record")
	_, isExist := emp["Emp 5"]
	fmt.Println(isExist)

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
