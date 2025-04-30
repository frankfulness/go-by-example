package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	// Arithmetic with arbitrary precision
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	// No type on numeric constant until given one, by conversion for ex;
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
