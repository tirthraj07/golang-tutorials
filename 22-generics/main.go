package main

import (
	"errors"
	"fmt"
	"strings"
)

func printSlice[T any](arr []T) {
	for _, val := range arr {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

func contains[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}

	return false
}

type Stack[T any] struct {
	elements []T
}

func (s Stack[T]) Top() (T, error) {
	if len(s.elements) == 0 {
		var zero T

		return zero, errors.New("stack is empty")
	}

	return s.elements[len(s.elements)-1], nil
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.elements) == 0 {
		var zero T

		return zero, errors.New("stack underflow")
	}

	lastIdx := len(s.elements) - 1

	lastElement := s.elements[lastIdx]

	s.elements = s.elements[:lastIdx]

	return lastElement, nil
}

func (s Stack[T]) String() string {
	if len(s.elements) == 0 {
		return ""
	}

	var builder strings.Builder

	for idx, element := range s.elements {
		builder.WriteString(fmt.Sprintf("%v", element))

		if idx != len(s.elements)-1 {
			builder.WriteString(" ")
		}
	}

	return builder.String()
}

func (s Stack[T]) Empty() bool {
	return len(s.elements) == 0
}

func main() {
	nums := []int{10, 20, 30}
	printSlice(nums)
	names := []string{"golang", "java", "c++"}
	printSlice(names)

	fmt.Println(contains([]int{1, 2, 3}, 2))
	fmt.Println(contains([]string{"a", "b"}, "b"))

	st := Stack[int]{}
	fmt.Println(st)
	st.Push(10)
	st.Push(20)
	st.Push(30)
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
	st.Push(40)
	fmt.Println(st)
	for !st.Empty() {
		top, err := st.Top()
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		fmt.Println(top)
		st.Pop()
	}

}
