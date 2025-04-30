package main

import "fmt"

func main() {
	// Most basic type with a single condition:
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// A classic initial/condition/after for loop:
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Another basic way to “do this N times” iteration is range over an integer
	for k := range 3 {
		fmt.Println("k range: ", k)
	}

	// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function:
	for {
		fmt.Println("infinite loop if no break or break case")
		break
	}

	// Continue to the next iteration of the loop:
	for n := range 8 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
