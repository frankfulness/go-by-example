package main

import (
	"fmt"
	"maps"
	// "maps"
)

func main() {
	/* Maps are Go’s built-in associative data type (sometimes called hashes or dicts in other languages).
	   To create an empty map, use the builtin make: make(map[key-type]val-type).
	*/
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 14

	// Printing a map with e.g. fmt.Println will show all of its key/value pairs.
	fmt.Println("Map: ", m)

	// Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// If the key doesn’t exist, the zero value of the value type is returned.
	v3 := m["k3"]
	fmt.Println("v3: ", v3)
	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len of m: ", len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("deleted key/value from m: ", m)

	// To remove all key/value pairs from a map, use the clear builtin.
	clear(m)
	fmt.Println("cleared map: ", m)

	/* The optional second return value when getting a value from a map indicates if the key was present in the map.
	This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	Here we didn’t need the value itself, so we ignored it with the blank identifier _. */
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	// Can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("new map: ", n)

	// The maps package contains a number of useful utility functions for maps.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

	n2["baz"] = 3
	if maps.Equal(n, n2) {
		fmt.Println("n =- n2 still")
	} else {
		fmt.Println("n != n2 as n2 now is: ", n2)
	}
}
