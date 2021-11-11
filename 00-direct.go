package main

import "fmt"

func DirectCallInlineFunction() {
	fmt.Println("Ok, lets do fibsies.")

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}

func DirectCall() {
	fmt.Println(fact(7))
}


func fact(n int) int {
	if n==0 {
		return 1
	}
	return n * fact(n-1)
}

// Recursive func calculating factorial of a positive integer
func factorialCalc(number int) int {
	// This is the base condition of number is 0 or 1,
	// the function will return 1.
	if number == 0 || number == 1 {
		return 1
	}

	// If negative arg it returns a message and -1.
	if number < 0 {
		fmt.Println("Invalid number")
		return -1
	}

	// Recursive call to self with decrement to argument by 1
	// allowing the return to eventually reach the base case.
	return number*factorialCalc(number-1)
}