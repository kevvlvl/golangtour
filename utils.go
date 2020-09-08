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