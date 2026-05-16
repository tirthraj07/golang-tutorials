package main

import (
	"fmt"
	"reflect"
	"time"
)

/*
A Struct is a collection of fields that are grouped together

Struct Memory Layout
Struct fields stored contiguously.
Order matters for:
	alignment
	padding
	memory usage

Go uses:
		tag based field ordering
		compiler optimizations
		memory layout guarantees
*/

type User struct {
	UserId    string
	Name      string
	Age       int
	Email     string
	createdAt time.Time
}

func updateByValue(user User) {
	user.Name = "Updated"
}

func updateByPointer(user *User) {
	user.Name = "Updated"
	/*
		This works
		even though: user is pointer

		Go automatically dereferences struct pointers for field access.
		Equivalent to conceptually - (*user).Name
	*/
}

// Nested Struct
type Address struct {
	City    string
	State   string
	Country string
}

type Order struct {
	OrderId         string
	ProductName     string
	DeliveryAddress Address
}

// Non-comparable Struct
type Engineer struct {
	Name   string
	Skills []string
}

func main() {

	// Basic Instantiation
	user := User{
		UserId:    "abc123",
		Name:      "Tirthraj Mahajan",
		Age:       21,
		Email:     "tirthraj@gmail.com",
		createdAt: time.Now(),
	}

	fmt.Println(user)

	/*
		-- VERY IMPORTANT --
		Structs are value types
		Meaning u1 := u2 creates a FULL COPY unlike Java Objects
	*/

	copyUser := user
	copyUser.Name = "Vardhan Dongre"
	fmt.Println(copyUser)

	user.Email = "tirthraj@google.com"
	fmt.Println(user)
	fmt.Println(copyUser)

	// You can also instantiate like this
	u2 := User{"abc234", "Harsh B", 21, "chotubadmash@gmail.com", time.Now()}
	fmt.Println(u2)

	/*
		This is generally discouraged
		Use Named Fields like
		User{
			Name: "Tirthraj",
			Age: 21,
		}
		as they are MUCH SAFER
		because:
			field order changes won't break silently
			more readable
			self-documenting
	*/

	// Structs in Functions

	// Pass By Value
	updateByValue(user)
	fmt.Println(user) // Still "Tithraj Mahajan". Caller unchanged.

	// Pass By Pointer
	updateByPointer(&user)
	fmt.Println(user) // Name "Updated". Caller modified.

	// Anonymous Structs
	person := struct {
		Name string
		Age  int
	}{
		Name: "Tirthraj",
		Age:  21,
	}
	fmt.Println(person)

	/*
		Useful for:
			temporary data
			testing
			JSON responses
	*/

	// Nested Structs
	order := Order{
		OrderId:     "123",
		ProductName: "Shampoo",
		DeliveryAddress: Address{
			City:    "Pune",
			State:   "Maharashtra",
			Country: "India",
		},
	}

	fmt.Println(order)
	fmt.Println(order.DeliveryAddress.Country)

	// Comparing Structs
	// Structs ARE comparable if all fields comparable.
	copyUser = user
	fmt.Println(user == copyUser) // true

	/*
		Note: A struct is comparable ONLY if: ALL its fields are comparable.
		If a Struct has Maps or Slices, the struct becomes: non-comparable
	*/
	e1 := Engineer{
		Name:   "Tirthraj",
		Skills: []string{"Go"},
	}

	e2 := Engineer{
		Name:   "Tirthraj",
		Skills: []string{"Go"},
	}

	fmt.Println(e1)
	fmt.Println(e2)
	// fmt.Println(e1 == e2)	// invalid operation: e1 == e2 (struct containing []string cannot be compared)
	// Go recursively checks comparability of all fields.

	/*
		Generally comparable:
			integers
			floats
			strings
			booleans
			pointers
			arrays (if elements comparable)
			structs (if all fields comparable)
		NOT Comparable
			slices
			maps
			functions
	*/

	// So How Do We Compare Structs With Slices?
	// Use: reflect.DeepEqual

	fmt.Println(reflect.DeepEqual(e1, e2)) // true

	/*
		reflect.DeepEqual:
			recursive comparison
			runtime reflection
			slower
			less type-safe

		Useful mostly for:
			testing
			debugging
			utilities

		Not always ideal in hot production paths.

		More on this at the end.
	*/

	// -----------
	/*
		Important Concept - Type Identity vs Type Equivalence
		Go defines:
			Type Identity - how it names things
			vs
			Type Equivalence - how it structures things

		Anonymous Structs don't have a formal name type like 'User' or 'Order' or 'Address'
		They are defined entirely by their shape (fields and type)

		In Go, two anonymous structs are considered exactly the same type if they
			- have the exact same fields
			- in the exact same order
			- with the exact same types

		For Named Structs - Go uses strict Nominal Typing for named types.
		Even though Struct1 and Struct2 have the exact same underlying memory structure (same fields, same order),
		Go treats them as completely isolated, distinct types.
		You cannot compare an apple to an orange, even if they weigh the same.

		Named to Anonymous
		Go allows you to assign or compare a named type with an unnamed type if their underlying structural layouts are identical.
	*/

	obj1 := struct {
		a string
		b int
	}{
		a: "abc",
		b: 10,
	}

	obj2 := struct {
		a string
		b int
	}{
		a: "abc",
		b: 10,
	}

	fmt.Println(obj1 == obj2) // true

	/*
		obj1 and obj2 are both anonymous (unnamed) structs
		Because they are the same type,
		and their underlying values ("abc" and 10) are identical,
		they are perfectly comparable and evaluate to true
	*/

	type Struct1 struct {
		a string
		b int
	}

	type Struct2 struct {
		a string
		b int
	}

	obj3 := Struct1{
		a: "abc",
		b: 10,
	}

	obj4 := Struct2{
		a: "abc",
		b: 10,
	}
	// fmt.Println(obj3 == obj4)	// invalid operation: obj3 == obj4 (mismatched types Struct1 and Struct2)
	fmt.Println(obj3 == Struct1(obj4)) // True

	/*
		So we need to explicity convert Struct2 object to Struct1 inorder for them to be comparable.
	*/

	obj5 := struct {
		a string
		b int
	}{
		a: "abc",
		b: 10,
	}
	fmt.Println(obj5 == obj3) // true

	/*
		Because obj5 has no formal name to conflict with Struct1,
		the Go compiler falls back to Structural Typing.

		It checks the shape,
		sees that they perfectly match,
		allows the comparison,
		and evaluates it to true based on their values.
	*/
}

/*
What reflect.DeepEqual Actually Does
It recursively walks values at RUNTIME.
Conceptually -
	compare struct
	→ compare each field
	→ recurse into slices/maps/structs
	→ recurse deeper

It uses relection to compare the values.
Meaning:
	runtime type inspection
	dynamic traversal


Why Reflection Is Slower
Normal comparison: a == b -> Compiler knows types at compile time. Very optimized.
But, reflect.DeepEqual(a, b)
must:
	inspect types dynamically
	recurse generically
	allocate sometimes
	use runtime metadata
Much more expensive.

For Hot Production Paths? Can become problematic.

Especially:
	large nested objects
	high-frequency comparisons
	low-latency systems

BUT Performance Is NOT The Biggest Issue
The BIGGER issue is: semantics

Nil Slice vs Empty Slice : reflect.DeepEqual(a, b) returns false. But Many Applications Consider Them Equivalent. Reflection cannot know your intent.
Reflection can compare fields you may not WANT included semantically. Example: createdAt, mutexes, caches, etc. DeepEqual blindly compares everything

Reflection bypasses compiler guarantees.
You lose:
	type-specific optimizations
	compile-time clarity
	discoverability

Reflection-heavy code often becomes:
	harder to debug
	harder to maintain
	harder to optimize

*/
