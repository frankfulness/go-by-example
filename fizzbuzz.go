package fizzbuzz

import "fmt"

func main() {
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
}
