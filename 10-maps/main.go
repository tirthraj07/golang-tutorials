package main

import (
	"fmt"
	"maps"
)

func main() {
	// Syntax : map[keyType]valueType

	// Nill Map
	var m1 map[string]string
	printMap(m1) // m == nil ? : true

	// Note: Writing to nill maps will cause panic
	// m1["abc"] = "xyz"	// panic: assignment to entry in nil map. Because internal hashmap not initialized.

	m2 := make(map[string]string)
	printMap(m2) // m == nil ? : false

	m := map[string]string{
		"name":  "Tirthraj Mahajan",
		"class": "BE",
		"div":   "02",
		"dept":  "Computer Engineering",
	}

	printMap(m)

	// Accessing value from a key
	val := m["name"]
	fmt.Println(val)

	// Accessing a non-existent value from map
	val = m["college"]
	fmt.Println(val)

	// Note: When a value does not exist, map returns zerovalue
	// What is zerovalue? it is default value of a type
	// int -> 0, bool -> false, string -> ""

	// How to know if a value exist in the map?
	// map returns another argument when we access it -> ok
	val, ok := m["college"]
	if !ok {
		fmt.Println("Value does not exist")
	} else {
		fmt.Println("Value exists!")
	}

	// Slighly better way (so that we don't declare ok constantly)
	if _, ok := m["college"]; !ok {
		fmt.Println("Value does not exist")
	} else {
		fmt.Println("Value exists!")
	}

	// Delete Keys
	// No error if key missing. Safe operation.
	delete(m, "div")
	printMap(m)

	// Iterating a map
	for key, value := range m {
		fmt.Println(key, value)
	}

	/*
		Important Note -
			Map iteration order is NOT guaranteed.
			Output order may change.
			This is intentional.
			Go deliberately randomizes iteration order.
		To prevent developers from accidentally depending on ordering. Huge design decision.
	*/

	// Maps are similar to slices m2 := m1 References same hashmap
	// Maps are NOT comparable. Only valid comparison: m == nil
	m3 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	m4 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Printf("m3 equals m4 ? %v\n", maps.Equal(m3, m4))

	// Nested Maps
	users := map[string]map[string]int{
		"Tirthraj": {
			"age": 21,
		},
	}
	fmt.Println(users)

	// Common Production Pattern — Frequency Counter
	words := []string{"go", "java", "go"}
	freq := make(map[string]int)

	for _, word := range words {
		freq[word]++
	}

	fmt.Println(freq)

	// Important = Maps are NOT safe for concurrent writes.
	/*
		This causes:
		fatal error: concurrent map writes
	*/
}

func printMap(m map[string]string) {
	fmt.Println("-------")
	fmt.Println(m)
	fmt.Printf("m == nil ? : %v\n", m == nil)
	fmt.Printf("len(m) ? %v\n", len(m)) // Returns number of key-value pairs.
	fmt.Println("-------")
}
