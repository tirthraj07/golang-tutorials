package main

import "fmt"

const tutorialName = "Constants" // Note: You cannot use shorthand expression here - view ls 03

func main() {
	const name string = "Tirthraj Mahajan"

	// name = "Ninad Palsule"
	// ERROR : cannot assign to name (neither addressable nor a map index expression)

	const age = 30

	// You can declare multiple constants simultaneously like this
	const (
		port = 8080
		host = "localhost"
	)

	fmt.Println(port, host)

}

/*

Unused constants are allowed in Go and do not cause compilation errors.
Unlike variables, unused constants (package-level or local) are ignored
by the compiler, do not take up memory space in the final executable,
and are not compiled into the binary, as they are resolved to their values at
compile time
*/
