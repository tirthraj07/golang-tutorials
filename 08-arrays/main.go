package main

import "fmt"

func main() {

	/*
		Arrays are important mainly because:
			slices are built on top of them
			understanding arrays helps understand memory
			they teach value semantics
		But in real-world Go:
			arrays are used rarely
			slices dominate almost everything
	*/

	// Basic Array
	var arr [5]int // Size is part of the type itself.
	fmt.Println(len(arr))
	fmt.Println(arr)
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)

	// Short Declaration
	names := [3]string{"Go", "Java", "Python"}
	fmt.Println(len(names))
	fmt.Println(names)

	// Arrays Are Value Types (Important)
	// Values types are copied entirely (Not like reference types)
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1 // Note: arr1 values are COPIED to arr2 (FULL COPY). This behaviour is different from Java
	arr2[0] = 100
	arr2[1] = 200
	fmt.Println(arr1) // [1 2 3]
	fmt.Println(arr2) // [100 200 3]

	// Arrays as functional Arguments
	numbers := [3]int{1, 2, 3}
	modify(numbers)      // array COPY passed into function
	fmt.Println(numbers) // [1 2 3]

	// Pointer to Array
	modifyArr(&numbers)  // Now original array changes.
	fmt.Println(numbers) // [999 2 3]

	// Iterating Arrays
	// Classic Loop
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	// Range Loop
	for index, value := range numbers {
		fmt.Println(index, value)
	}
	// Ignore index
	for _, value := range numbers {
		fmt.Println(value)
	}
	// Ignore value
	for index := range numbers {
		fmt.Println(index)
	}

	// You can also declare arrays like these
	nums := [...]int{1, 2, 3, 4} // Compiler infers size automatically.
	fmt.Println(nums)

	// Multi-Dimensional Arrays
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(matrix)
}

func modify(arr [3]int) {
	arr[0] = 999
	fmt.Println(arr) // [999 2 3]
}

func modifyArr(arr *[3]int) {
	arr[0] = 999
}
