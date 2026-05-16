package main

import "fmt"

/*
Struct embedding is one of the MOST important Go concepts.
Because this is Go’s replacement for: Inheritance

Embedding = Composition + Promotion

Go does NOT have:
	classes
	extends
	inheritance hierarchies

Instead Go prefers:
composition over inheritance
And struct embedding is a major part of that philosophy.

Java often uses:
	inheritance for reuse

Go prefers:
	composition for reuse

Go embedding makes composition ergonomic enough that: inheritance becomes unnecessary most of the time.
This is one of Go’s most elegant ideas.
*/

type Address struct {
	City string
}

type User struct {
	Name string
	Address
}

/*
Notice: Address does NOT have a field name. This is embedding.
*/

type A struct {
	Name string
}

type B struct {
	Name string
}

type C struct {
	A
	B
}

func main() {
	user := User{
		Name: "Tirthraj Mahajan",
		Address: Address{
			City: "Pune",
		},
	}

	fmt.Println(user)
	fmt.Println(user.City)

	// User does NOT directly define: user.City, yet it works. Why? Because embedded fields/methods get promoted, meaning outer struct exposes them directly.'
	// Go internally treats: user.City like user.Address.City

	// This Is NOT Inheritance. VERY IMPORTANT.
	/*
		Embedding:
			does NOT create subtype relationship
			does NOT create polymorphic hierarchy
			does NOT create “is-a” relationship

		It creates:
			“has-a” relationship

		User HAS an Address
		NOT:
			User IS an Address
		Huge conceptual distinction.
	*/

	// Field Name Conflicts
	c := C{
		A: A{
			Name: "TJM",
		},
		B: B{
			Name: "ABC",
		},
	}
	// fmt.Println(c.Name)	// ambiguous selector c.Name
	// Must specify:
	fmt.Println(c.A.Name)
	fmt.Println(c.B.Name)

}
