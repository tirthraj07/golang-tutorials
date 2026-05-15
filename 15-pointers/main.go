package main

import "fmt"

func fooValue(x int) {
	x = 10
	fmt.Printf("Value in foo function : %v\n", x)
}

func fooPointer(x *int) {
	*x = 10
	fmt.Printf("Value in foo function : %v\n", *x)
}

func create() *int {
	x := 10

	return &x
}

func main() {
	x := 10

	ptr := &x

	fmt.Println(ptr)  // address-of operator
	fmt.Println(*ptr) // Deferencing

	// Modify Through Pointer
	*ptr = 50
	fmt.Println(x)

	/*
		It is important to note that - Go is always pass by value. Similar to java
		Thereform when we do this
	*/
	x = 20
	fooValue(x)    // Inside this -> x = 10
	fmt.Println(x) // Outside -> x = 20

	// If we want to modify the value, we must pass in the reference
	fooPointer(&x) // Inside this -> x = 10
	fmt.Println(x) // Outside -> x = 10

	// Nil Pointers
	var nilptr *int
	fmt.Println(nilptr) // <nil>
	// fmt.Println(*nilptr)	// Dereferencing Nil Pointer PANICS -> invalid memory address or nil pointer dereference

	// new()
	ptr = new(int)
	*ptr = 10
	fmt.Println(ptr)
	fmt.Println(*ptr)

	// Escape Analysis (VERY IMPORTANT)
	ptr = create()
	/*
		This looks dangerous.
		Because:
			x local variable
			function exits

		Shouldn’t pointer become invalid?
		In C/C++ → dangerous bug
		But Go compiler detects: x escapes function
		So: allocates x on heap

		Safe.

		Go automatically manages memory lifetime.

		No dangling pointers.
	*/

}
