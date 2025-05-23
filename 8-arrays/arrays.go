package main

import "fmt"

func main() {
	// Create an array that will hold 5 ints
	// The type of elements and length are both part of the array’s type. By default an array is zero-valued, which for ints means 0s.
	var a [5]int
	fmt.Println("emp: ", a) // output should be [0 0 0 0 0]

	// Set a value at an index using the array[index] = value syntax, and get a value with array[index].
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("Length via len : ", len(a)) // The builtin len returns the length of an array.

	// Declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	// Have the compiler count the number of elements for me with ...
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	// Specify the index with :, the elements in between will be zeroed.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx :", b)

	// Array types are one-dimensional, but we can compose types to build multi-dimensional data structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Create and initialize multi-dimensional arrays at once too.
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
