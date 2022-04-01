package main

import (
	"fmt"

	greeting "./greetings"
	hello "./hello"
)

func main() {
	var intVariable int // Declaration
	intVariable = 45245 // Initialization
	fmt.Println(intVariable)

	var myVariable = "Hello, Variables" // Declaration wit Initialization
	fmt.Println(myVariable)

	//Calling hello method from greetings folder
	greeting.Hello("greetings.go\\greetings")

	//Calling hello method from hello
	hello.PrintHello()
}
