package main

import "fmt"

func apaajalah()  {
	type Person struct { // Membuat struct bernama Person
		Name string
		Age int
		Money float64
	}
	p1 := Person{Name: "RAFI", Age: 19, Money: 1000000}
	p2 := Person{Name: "SLEEK", Age: 21, Money: 2000000}

	fmt.Println(p1, p2) // Output: {RAFI 19 1000000} {SLEEK 21 2000000}
}