package main

import "fmt"

// Go supports recursive functions, which are functions that run themselves until a base case is met.
// This fact function calls itself until it reaches the base case of fact(0).
func fact(n int) int {
	fmt.Println("n: ", n)
	// base case:
	if n == 0 {
		fmt.Println("Base case met, exit.")
		return 1
	}
	// factors and invokes recursive function with new argument
	return n * fact(n-1)
}

func main() {
	/* Anonymous functions can also be recursive, but this requires explicitly declaring
	a variable with var to store the function before it’s defined */
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		// Base case
		if n < 2 {
			return n
		}
		// recursive logic
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
