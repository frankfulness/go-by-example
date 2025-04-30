package main

import "fmt"

func main() {
	// FizzBuzz to show if, else if, else with && logical operators within a for range loop
	for i := range 100 {
		// divisible by 3 and 5 must go first, otherwise program will execute Fizz or Buzz
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	for j := range 10 {
		// demonstrate || logical operator too
		if j%2 == 0 || j%3 == 0 {
			fmt.Println(j, "is divisible by 2 or 3")
		} else {
			fmt.Println(j, " is not divisible")
		}
	}

	// A statement can precede conditionals; any variables declared in this statement are available in the current and all subsequent branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
