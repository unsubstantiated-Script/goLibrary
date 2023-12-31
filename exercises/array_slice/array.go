package array_slice

import "fmt"

func RollArrays() {
	// Declaring an array
	var ni [7]int
	// Assign a value to index position zero.
	ni[0] = 42
	fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

	// Array literal
	ni2 := [4]int{55, 56, 57, 58}
	fmt.Printf("%#v \t\t\t\t %T\n", ni2, ni2)

	// Array literal
	// Compiler counts elements
	ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
	fmt.Printf("%#v \t %T\n", ns, ns)

	// Use the built in len function
	fmt.Println(len(ni))
	fmt.Println(len(ni2))
	fmt.Println(len(ns))
}
