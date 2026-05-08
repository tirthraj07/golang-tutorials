package main

import (
	"fmt"
	"strings"
)

/*
Strings in Golang

Strings look simple but they contain -
	memory concepts
	immutability
	Unicode
	UTF-8 encoding
	byte vs rune distinction

Strings in Go is designed around UTF-8.
*/

func main() {

	// Basic Strings
	name := "Tirthraj"
	fmt.Println(name)

	// Important Concept - Strings Are Immutable
	// You CANNOT modify string characters directly.
	// s := "Go"
	// s[0] = 'N'	// cannot assign to s[0] (neither addressable nor a map index expression)

	/*
		Why are they immutable?
		Internally strings are read-only byte sequences
		Conceptually:
			type string struct {
				ptr *byte
				len int
			}
		Immutable by design.

		Benefits of Immutability
			safer sharing
			easier concurrency
			fewer bugs
			better optimizations
	*/

	// String Indexing
	str := "Golang"
	fmt.Println(str[0]) // Output - 71

	/*
		Why 71?
		Because indexing returns BYTE value.
		ASCII:
			'G' = 71
	*/

	// Convert to Character
	fmt.Println(string(str[0])) // Output - 'G'

	// Important concept - STRINGS ARE BYTES
	// Go strings are UTF-8 encoded byte sequences.
	str = "hello"
	fmt.Println(len(str)) // Output - 5
	/*
		Why 5? because each ascii character in "hello" = 1 byte
	*/

	// Unicode Problem
	str = "नमस्ते"
	fmt.Println(str)
	fmt.Println(len(str)) // Output - 18

	/*
		Why 18?
		Because:
			UTF-8 uses variable-width encoding
			many Unicode chars use multiple bytes

		This is one of the biggest confusions in Golang
		len(string) returns the number of BYTES not CHARACTERS
	*/

	// Important Concept - Rune
	// Go uses 'rune' for Unicode code points
	// rune is alias for int32

	// Convert String to Rune Slice
	runes := []rune(str)
	fmt.Println(runes)      // [2344 2350 2360 2381 2340 2375]
	fmt.Println(len(runes)) // output - 6
	// Now length becomes actual character count.

	// Iterating Strings Properly
	/*
		WRONG for Unicode
		for i := 0; i < len(str); i++ {
			fmt.Println(string(str[i]))
		}
		Breaks Unicode chars.
	*/

	// Correct Approach - use range
	for index, char := range str {
		fmt.Println(index, char, string(char))
	}

	/*
		Why range works?
			Range automatically decodes UTF-8 into runes.
	*/

	// String Concatenation (Similar to Java)
	first := "Hello"
	second := "World"
	result := first + " " + second
	fmt.Println(result)

	/*
		Important Performance Insight
		Repeated string concatenation is expensive.
		Because:
			strings immutable
			new string allocated each time
	*/

	// Use strings.Builder for efficient string construction
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" World")
	result = builder.String()
	fmt.Println(result)

	/*
		Used heavily in:
		parsers
		serializers
	*/

	// Multi-Line Strings
	text := `
	Hello
	World
	`
	fmt.Println(text)
	// Called raw string literals.
	// Useful for: SQL JSON regex templates

	// Important Difference - Double Quotes vs Backticks
	// Double Quotes supports escape characters
	// Backticks does not support escape characters
	a := "Hello\nWorld"
	b := `Hello\nWorld`
	fmt.Println(a, b)

	// Common String Functions (strings package)
	str = "Hello World"
	fmt.Println(strings.Contains(str, "Hello"))
	fmt.Println(strings.Count(str, "l"))
	fmt.Println(strings.Index(str, "World"))
	fmt.Println(strings.LastIndex(str, "World"))
	fmt.Println(strings.Replace(str, "World", "Go", 1))
	fmt.Println(strings.ReplaceAll(str, "World", "Go"))
	fmt.Println(strings.Split(str, " "))
	fmt.Println(strings.Join(strings.Split(str, " "), "-"))
	fmt.Println(strings.TrimSpace(str))

	// Strings vs Byte Slices (Very important)
	// String is immutable UTF-8 Bytes
	// byte slice ([]byte) is mutable

	// Conversion
	str = "hello"
	bytes := []byte(str)
	bytes[0] = 'H'
	fmt.Println(string(b)) // Hello

	/*
		Note:
			High-performance Go systems often:
				avoid excessive string allocations
				use byte slices directly
				delay string conversion

			Very important in:
				networking
				parsers
				databases
				Kafka clients
				HTTP servers
	*/

	// Note - Strings are COMPARABLE (not like java)
	a = "hello"
	b = "hello"
	fmt.Printf("a == b ? %v\n", a == b)

	// Lexicographical comparison su pported too ("a" < "b")

}
