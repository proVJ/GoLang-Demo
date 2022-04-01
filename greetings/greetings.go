package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// GLOBAL LEVEL VARIABLE (pascal case)
// this variable will accessible for all packages
var GlobalVariable = "GLOBAL LEVE VARIABLE"

// PACKAGE LEVE VARIABLE (calmel case)
// this variable will accessible to this package only
var packageLevelVariable int = 8784524
