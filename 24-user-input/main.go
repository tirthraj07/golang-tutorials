package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// --- 1. Integer Array ---
	fmt.Print("Enter the size of the integer array: ")
	scanner.Scan() // Reads the whole first line

	// Convert the input string to an integer
	intSize, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	fmt.Printf("Enter %d integers (on ONE line, separated by spaces): ", intSize)
	scanner.Scan() // Reads the entire second line (e.g., "10 20 30")
	intLine := scanner.Text()

	// Split the line into a slice of strings based on spaces
	// strings.Fields safely handles multiple spaces between words
	strNumbers := strings.Fields(intLine)

	// Create our integer array and convert each string to an int
	intArray := make([]int, 0, intSize)
	for i := 0; i < intSize && i < len(strNumbers); i++ {
		num, _ := strconv.Atoi(strNumbers[i])
		intArray = append(intArray, num)
	}

	// --- 2. String Array ---
	fmt.Print("Enter the size of the string array: ")
	scanner.Scan()
	strSize, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	fmt.Printf("Enter %d strings (on ONE line, separated by spaces): ", strSize)
	scanner.Scan()
	strLine := scanner.Text()

	// Split the string line
	strArray := strings.Fields(strLine)

	// --- Print Results ---
	fmt.Println("\n--- Final Output ---")
	fmt.Printf("Integer Array: %v\n", intArray)
	fmt.Printf("String Array:  %v\n", strArray)
}
