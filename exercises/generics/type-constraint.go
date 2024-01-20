package generics

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

// THis generic will take in and return an int or a float
// Type Constraint
func addT[T int | float64](a, b T) T {
	return a + b
}

// Type set
type myNumbers interface {
	int | float64
}

// Used here in square brackets
func addS[T myNumbers](a, b T) T {
	return a + b
}

func RollTypeConstraint() {
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))

	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.2, 2.2))

	fmt.Println(addS(1, 2))
	fmt.Println(addS(1.2, 2.2))
}
