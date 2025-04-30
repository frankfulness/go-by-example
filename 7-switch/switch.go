package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic switch
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One.")
	case 2:
		fmt.Println("Two..")
	case 3:
		fmt.Println("Three!")
	}

	// Commas to separate multiple expressions in the same case statement. Use optional default case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Work, it's still a weekday.")
	}

	// Switch without an expression is an alternate way to express if/else logic.
	// Show how how the case expressions can be non-constants.
	t := time.Now()
	switch { // Note no expression
	case t.Hour() < 12:
		fmt.Println("Before Noon.")
	default:
		fmt.Println("After Noon.")
	}

	/* A type switch compares types instead of values.
	Use this to discover the type of an interface value.
	In this example, the variable t will have the type corresponding to its clause. */
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("You are a bool.")
		case int:
			fmt.Println("You are an int.")
		case string:
			fmt.Println("You are a string.")
		default:
			fmt.Printf("You are type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(33)
	whatAmI("Frank")
	whatAmI(10.0) // Float should be differrent type than int and hit default case.
}
