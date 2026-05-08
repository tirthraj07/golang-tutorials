package main

import "fmt"

func main() {

	// for is the only keyword used for looping

	// Classic for loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// While Loop using for loop
	i := 0
	for i < 3 {
		i += 2
	}
	fmt.Println(i)

	// Infinite loop
	// for {
	// 	fmt.Println("Tirthraj Mahajan")
	// }

	// Range
	i = 0
	// Note : Range has end excluded
	for i = range 10 {
		fmt.Println(i)
	}

	// Note: You can use continue and break as we do in other programming languages

}
