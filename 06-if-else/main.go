package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 13 {
		fmt.Println("You are a teenager")
	} else if age >= 3 {
		fmt.Println("You are a child")
	} else {
		fmt.Println("You are a baby")
	}

	/*
		Go has a special pattern:
		if initialization; condition {
			// code to execute
		}
	*/

	if role := "user"; role == "admin" {
		fmt.Println("Permission Granted")
	} else {
		fmt.Printf("Permission not granted. Current role - %v : Required role - admin\n", role)
	}

	/*
		The variable exists ONLY inside the if block
		This is heavily used in real Go code.
		Real Production Example
		if err := validateUser(); err != nil {
			fmt.Println(err)
			return
		}
		You will see this everywhere in Go.
		This is one of the most important idioms in the language.
	*/

	// Variable Shadowing (IMPORTANT)
	// This is a massive real-world Go issue.
	yoe := 5

	// The inner age shadows the outer one.
	if yoe := 10; yoe >= 10 {
		fmt.Println(yoe) // Prints 10
	}

	fmt.Println(yoe) // Prints 5

	// Note: Go Does NOT Have Ternary Operator
	// Go prioritizes readability over compactness.

}
