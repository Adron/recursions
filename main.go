package main

import "fmt"

func main() {
	DirectCall()
	DirectCallInlineFunction()



	// passing 0 as a parameter
	answer1 := factorialCalc(0)
	fmt.Println(answer1)

	// passing a positive integer
	answer2 := factorialCalc(5)
	fmt.Println(answer2)

	// passing a negative gives us the error message & -1
	answer3 := factorialCalc(-1)
	fmt.Println(answer3)

	// passing a positive integer
	answer4 := factorialCalc(10)
	fmt.Println(answer4)


	Indirect()

	// passing a positive
	// parameter which prints all
	// numbers from 1 - 10
	printOne(10)

	// this will not print
	// anything as it does not
	// follow the base case
	printOne(-1)
}
