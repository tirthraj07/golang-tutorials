package main

import "fmt"

/*
n Go: EVERYTHING has a type

Including:
	primitives
	structs
	interfaces
	functions
	slices
*/

type Speaker interface {
	Speak()
}

type Dog struct{}
type Cat struct{}

// Dog implements Speaker
func (d Dog) Speak() {
	fmt.Println("Woof Woof!")
}

// Cat implements Speaker
func (c Cat) Speak() {
	fmt.Println("Meoww Meoww")
}

func main() {
	var x any
	x = "Hello"

	num, ok := x.(int)
	if !ok {
		fmt.Println("Not Integer")
	} else {
		fmt.Println("Integer: ", num)
	}

	switch v := x.(type) {
	case int:
		fmt.Println("int: ", v)
	case string:
		fmt.Println("string: ", v)
	case bool:
		fmt.Println("boolean: ", v)
	default:
		fmt.Println("unknown")
	}

	var animal any = Dog{}

	_, ok = animal.(Speaker)
	fmt.Println(ok)
}
