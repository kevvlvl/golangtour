package main

import "fmt"

// Simple function that returns an int type
func add(x int, y int) int {
	return x + y
}

// Print an array length, capacity and contents alongside the message/log
func printArrayDetails(message string, numbers[] int) {
	fmt.Println(message, "Length: ", len(numbers), " - Capacity: ", cap(numbers), " Contents: ", numbers)
}

// iterate a slice and show its contents
func iterateSlice(numbers[] int) {

	// to omit i or v in the for loop, replace variable by underscore _
	for i, v := range numbers {
		fmt.Println("Current index: ", i, ", value: ", v)
	}
}