package main

import "fmt"

// Go’s structs are typed collections of fields.
// They’re useful for grouping data together to form records.

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 33
	return &p
}

func main() {
	// Create a new struct
	fmt.Println("Person: ", person{"Frank", 33})
	// Name the fields when initializing a struct.
	fmt.Println("Person: ", person{name: "Kristen", age: 32})

	// Omitted fields will be zero-valued.
	fmt.Println("Person: ", person{name: "Hope"})

	// An & prefix yields a pointer to the struct.
	fmt.Println("Person: ", &person{name: "Elijah", age: 4})

	// It’s idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Ezekiel"))

	// Access struct fields with a dot.
	s := person{name: "dad", age: 81}
	fmt.Println("I love my", s.name)

	// Also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 82
	fmt.Println(sp.age)

	/* If a struct type is only used for a single value,
	we don’t have to give it a name.
	The value can have an anonymous struct type.
	This technique is commonly used for table-driven tests. */
	dog := struct {
		name   string
		isGood bool
	}{
		"Goldie",
		true,
	}
	fmt.Println(dog)
}
