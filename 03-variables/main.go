package main

import "fmt"

func main() {
	// --- 1. String ---
	// Explict
	var name1 string = "golang"
	fmt.Println(name1)

	// Infer
	var name2 = "java"
	fmt.Println(name2)

	// Shorthand
	name3 := "python"
	fmt.Println(name3)

	// Note - In Go, short variable declarations (:=) are restricted to function bodies primarily to simplify language parsing
	// Since shorthand is not allowed, you must use the var keyword for package-level variables

	// --- 2. Boolean ---
	// Explict
	var isVerified1 bool = true
	fmt.Println(isVerified1)

	// Infer
	var isVerified2 = false
	fmt.Println(isVerified2)

	// Shorthand
	isVerified3 := true
	fmt.Println(isVerified3)

	// --- 3. Integer ---
	// Explict
	var age1 int = 21
	fmt.Println(age1)

	// Infer
	var age2 = 21
	fmt.Println(age2)

	// Shorthand
	age3 := 21
	fmt.Println(age3)

	// --- 4. Float ---
	// Explict
	var price1 float64 = 10.45
	fmt.Println(price1)

	// Infer
	var price2 = 23.211
	fmt.Println(price2)

	// Shorthand
	price3 := 34.2
	fmt.Println(price3)

	// Define First, Initialize Later -> You cannot use shorthand
	var address string
	address = "Pune, India"
	fmt.Println(address)
}
