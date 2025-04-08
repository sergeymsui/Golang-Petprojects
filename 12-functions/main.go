package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name  string
	fname string
	age   int
}

func main() {

	persons := []Person{
		{"Alice", "Mat", 11},
		{"Den", "Low", 5},
		{"Paul", "Pet", 20},
	}

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].age < persons[j].age
	})

	for _, person := range persons {
		fmt.Printf("name: %s, fname: %s, age: %d\n", person.name, person.fname, person.age)
	}
}
