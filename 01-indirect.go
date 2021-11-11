package main

import "fmt"

func Indirect() {
	fmt.Println("Indirect called.")
}

// recursive function for
// printing all numbers
// upto the number n
func printOne(n int) {
	// if the number is positive
	// print the number and call
	// the second function
	if n >= 0 {
		fmt.Println("In first function:", n)
		// call to the second function
		// which calls this first
		// function indirectly
		printTwo(n - 1)
	}
}

func printTwo(n int) {
	// if the number is positive
	// print the number and call
	// the second function
	if n >= 0 {
		fmt.Println("In second function:", n)
		// call to the first function
		printOne(n - 1)
	}
}