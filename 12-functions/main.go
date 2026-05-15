package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"
)

// Basic Function
func greet(name string) {
	fmt.Printf("Hello %v! Welcome to Golang Tutorial.\n", name)
}

// func add(a int, b int) vs func add(a, b int)
func add(a, b int) int {
	return a + b
}

// Multiple Return Value
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// Errors
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

// Recursion
func factorial(a int) (int, error) {
	if a < 0 {
		return 0, errors.New("cannot find factorial of negative numbers")
	}

	if a == 0 {
		return 1, nil
	}

	res, err := factorial(a - 1)

	if err != nil {
		return 0, err
	}

	return a * res, nil
}

// Random String
func generateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)

	if err != nil {
		return "", fmt.Errorf("unable to generate random string: %w", err)
	}

	result := make([]byte, length)

	for i, b := range randomBytes {
		result[i] = charset[int(b)%len(charset)]
	}

	return string(result), nil
}

// Wapper function
func wrapper(fn func(int) (string, error), arg int) (string, error) {
	start := time.Now()
	res, err := fn(arg)
	end := time.Now()
	fmt.Printf("Execution Time: %v\n", end.Sub(start))
	return res, err
}

func main() {
	// Basic Function and Function Call
	greet("Tirthraj")

	// Syntax Breakdown
	// func functionName(parameters) returnType
	res := add(1, 2)
	fmt.Println(res)

	// Multiple Return Values
	quotient, remainder := divide(10, 3)
	fmt.Printf("Quotient : %v\nRemainder : %v\n", quotient, remainder)

	/*
		WHY This Matters
		Go uses multiple returns heavily for:
			errors
			status handling
			parsing
			lookup operations
	*/

	// Huge Example
	// result, err := doSomething()
	// This is THE core Go error handling pattern.
	res, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Result: %v\n", res)
	}

	// Recursion
	fac, err := factorial(-1)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Result: %v\n", fac)
	}

	// Anonymous Function
	square := func(x int) int {
		return x * x
	}
	fmt.Println(square(5))

	// Random String generation function
	randomString, err := generateRandomString(10)
	if err != nil {
		fmt.Println("Unable to generate random string. Error: ", err)
	} else {
		fmt.Printf("Random String is: %v\n", randomString)
	}

	// Wrapper Functions
	randomString, err = wrapper(generateRandomString, 10000)
	if err != nil {
		fmt.Println("Unable to generate random string. Error: ", err)
	} else {
		fmt.Printf("Random String is: %v\n", randomString)
	}
}
