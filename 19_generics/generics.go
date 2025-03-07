package main

import (
	"fmt"
)

func printSlice[T comparable](items []T) { // comparable includes all types
	for _, item := range items {
		fmt.Println(item)
	}
}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//lifo
// type stack struct {
// 	elements []int
// }

func main() {

	// 	myStack := stack{
	// 		elements: []int{1, 2, 3},
	// 	}

	// 	fmt.Println(myStack)
	//printSlice(items [] int)
	// names := []string{"golang", "typescript"}
	// printStringSlice(names)
	//printSlice(names)
	vals := []bool{true, false, true}
	printSlice(vals)

}
