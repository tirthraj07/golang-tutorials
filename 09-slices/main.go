package main

import (
	"fmt"
	"slices"
)

/*
A slice is a dynamic window over an underlying array
Not the array itself.

Array ([5]int) -> Fixed-size storage.
Slice ([]int) -> A lightweight descriptor pointing to an array.

Internally a slice contains:
	pointer -> underlying array
	length
	capacity
*/

func printSliceDetails(arr []int) {
	fmt.Println("---------")
	fmt.Printf("Value of slice: %v\n", arr)
	fmt.Printf("arr == nil : %v\n", arr == nil)
	fmt.Printf("Length of arr : %v\n", len(arr))
	fmt.Printf("Capacity of arr: %v\n", cap(arr))
	fmt.Println("---------")
}

func main() {

	// Creating a basic slice
	// Declaring a slice
	var arr []int
	printSliceDetails(arr)
	// Value of slice: []
	// arr == nil : true
	// Length of arr : 0
	// Capacity of arr: 0

	// Initialize array
	// Type, initial length, initial capacity
	arr = make([]int, 0, 5)
	printSliceDetails(arr)
	// Value of slice: []
	// arr == nil : false
	// Length of arr : 0
	// Capacity of arr: 5

	/*
		Length - Number of accessible elements.
		Capacity - Maximum growth before reallocation required.
	*/

	// append - Important Concept
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	arr = append(arr, 4)
	printSliceDetails(arr)
	// Value of slice: [1 2 3 4]
	// arr == nil : false
	// Length of arr : 4
	// Capacity of arr: 5
	/*
		append returns a NEW slice value.
		So you must capture it.

		when you do: append(numbers, 1)
		Go may create:
			a new underlying array
			a new slice descriptor
		Therefore:
			your old slice variable may become outdated

		THIS IS THE KEY IDEA append() does NOT mutate the slice variable itself. It returns an updated slice.

	*/

	// Important
	// Nil slices are valid slices.
	// You can append to them safely.
	var numbers []int
	printSliceDetails(numbers)
	// ---------
	// Value of slice: []
	// arr == nil : true
	// Length of arr : 0
	// Capacity of arr: 0
	// ---------
	numbers = append(numbers, 1)
	printSliceDetails(numbers)
	// ---------
	// Value of slice: [1]
	// arr == nil : false
	// Length of arr : 1
	// Capacity of arr: 1
	// ---------

	// Shorthand declaration
	nums := []int{1, 2, 3, 4, 5}
	printSliceDetails(nums)
	// 	---------
	// Value of slice: [1 2 3 4 5]
	// arr == nil : false
	// Length of arr : 5
	// Capacity of arr: 5
	// ---------

	// Note: Since slice is lightweight descriptor pointing to array
	// slice2 := slice1 DOES NOT COPY THE ARRAY
	// Both share same underlying array.
	slice1 := []int{1, 2, 3}
	slice2 := slice1
	slice2[0] = 2
	printSliceDetails(slice1) // Value of slice: [2 2 3]
	printSliceDetails(slice2) // Value of slice: [2 2 3]
	// Changes in one changed both!

	/*
		This is where slices can get tricky
		Important concept to understand
			slices are pointers to arrays.
			append might create a new array
		Take this example
	*/
	slice3 := []int{1, 2, 3}
	slice4 := slice3
	printSliceDetails(slice3) // Capacity of arr: 3
	printSliceDetails(slice4) // Capacity of arr: 3

	/*
		Currently this is what is happening
		slice3 ----\
					-> [1,2,3]
		slice4 ----/
	*/

	// Now if we do
	slice4 = append(slice4, 4)
	printSliceDetails(slice4) // Capacity of arr: 6
	/*
		Since capacity of array was already reached before (3)
		When go tried to append 4, it created a new array of capacity 6
		and now slice4 is pointing to THAT NEW ARRAY

		therefore, slice3 and slice4 are now pointing to two separate arrays
	*/
	printSliceDetails(slice3) // [1 2 3]
	printSliceDetails(slice4) // [1 2 3 4]

	// So this is where slices can get buggy - Do NOT rely on append side effects.

	// How to copy slices then?
	// If you need independent data: use copy() function
	original := []int{1, 2, 3, 4, 5}
	copySlice := make([]int, len(original))
	copy(copySlice, original)
	// Copy slice are now independent memory.
	printSliceDetails(original)
	printSliceDetails(copySlice)

	// Slices are NOT comparable.
	// a == b : INVALID
	// Slices can only be compared to nil
	// Why? Because slices are descriptors over mutable shared memory. Comparison semantics become complicated.

	// important concept - Range Over Slices
	vec := []int{1, 2, 3, 4, 5, 6, 7, 9, 9}
	for index, value := range vec {
		fmt.Println(index, value)
	}

	// Gotcha - Range copies value
	for _, value := range vec {
		value += 10
	}

	fmt.Println(vec) // values didn't change

	// You need index mutation:
	for i := range vec {
		vec[i] += 10
	}

	fmt.Println(vec) // values changed

	// Important - Slicing a Slice
	a := []int{1, 2, 3, 4, 5}
	b := a[1:4]
	c := b[1:2]
	// All may share SAME underlying array
	// This becomes extremely important for memory leaks.
	// c makes the entire a array alive in memory (really bad is you no longer required a and it was very big)
	// Fix: Copy needed data into new slice.
	fmt.Println(c)

	// Reversing a slice
	x := []int{1, 2, 3, 4, 5, 6}
	y := make([]int, len(x))
	for idx := range x {
		y[len(x)-idx-1] = x[idx]
	}
	fmt.Println(x)
	fmt.Println(y)

	// in place reverse
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		temp := x[i]
		x[i] = x[j]
		x[j] = temp
	}
	fmt.Println(x)

	x = removeAt(x, 2)
	fmt.Println(x)

	// Comparing two slices
	p := []int{1, 2, 3}
	q := []int{1, 2, 3}
	fmt.Println(slices.Equal(p, q))
	slices.Reverse(x)
	fmt.Println(x)
	slices.Sort(x)
	fmt.Println(x)
	fmt.Println(slices.Max(x))
	fmt.Println(slices.Min(x))
	fmt.Println(slices.Contains(x, 3))
	fmt.Println(slices.Index(x, 3))

}

func removeAt(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
