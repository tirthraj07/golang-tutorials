package main

import "fmt"

func main() {
	// Basic Switch

	day := "Thursday"

	switch day {
	case "Monday":
		fmt.Println("Start of work week")
	case "Friday":
		fmt.Println("Weekend Soon")
	case "Sunday", "Saturday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Normal Day")
	}

	// break is automatic
	// cases do NOT fall through by default

	// Go does support explicit fallthrough.
	num := 1
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
		fallthrough
	default:
		fmt.Println("Greater than Three")
	}

	// Switch Without Expression (Important)
	/*
		This is where Go becomes interesting.
		You can write:
		switch {
		case condition1:
		case condition2:
		}
	*/

	score := 82
	switch {
	case score >= 90:
		fmt.Println("Grade A")

	case score >= 75:
		fmt.Println("Grade B")

	case score >= 50:
		fmt.Println("Grade C")

	default:
		fmt.Println("Fail")
	}

	// This is basically a cleaner if-else-if

	// Scoped Switch
	switch num := 10; {
	case num > 5:
		fmt.Println("Greater than 5") // num exists only inside switch.
	}

	// Type Switch
	/*
		This becomes critical when learning interfaces
		switch value := x.(type) {
		case int:
		case string:
		}

		You’ll use this later for:
			1. interfaces
			2. polymorphism
			3. serializers
			4. middleware
			5. generic systems
	*/

}
