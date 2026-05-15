package main

import "fmt"

/*
A function can REMEMBER variables from where it was created.
That remembered environment is called: Closure
*/

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func logger(serviceName string) func(string) {
	return func(message string) {
		fmt.Printf("[%v] %v\n", serviceName, message)
	}
}

func main() {

	// Take this simple example -
	name := "Tirthraj"

	greet := func() {
		fmt.Printf("Hello %v\n", name)
	}

	greet()

	// Why Is This Interesting?
	// Because: greet does NOT define name
	// Yet it can access it.
	// Why? Because it “closed over” the surrounding scope. Hence: closure

	// Another amazing example
	c := counter()

	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 3

	/*
		THIS FEELS WEIRD Initially
		Because: counter() already finished execution.
		So why does: count exist?

		Answer: Because returned function captured it. Go keeps it alive.


		When we do `c := counter()`
		Go creates something conceptually like -
		'''
			function
			+
			hidden state:
				count = 0
		'''

		// Then every call c() modifies SAME captured variable.
	*/

	// Example
	authLogger := logger("AUTH")
	authLogger("Login successful") // [AUTH] Login successful
	authLogger("Token expired")    // [AUTH] Token expired
}

/*
Garbage collection happens when: an object is no longer reachable by the program

In closures - 'count' is still reachable through returned function. So GC does NOT collect it.

What Normally Happens?

Normally local variables live on: stack
When function exits:
	stack frame destroyed
	variables disappear

But Closures Change Everything
Here:
```
	return func() int {
		count++
		return count
	}
```

returned function STILL needs: count
after counter() exits.
So Go CANNOT destroy it.

Therefore Go performs: escape analysis
Compiler realizes: count escapes function lifetime

So instead of stack: allocate on HEAP

Instead of:
stack:
count

Go moves it to:
heap:
count

Returned closure points to it.

Garbage Collection Happens When Closure Dies - After 'c' becomes unreachable and no references remain
*/
