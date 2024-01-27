package generics

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

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

// Type set "~" means include all values of that type as well as underlying values
type myNumbers interface {
	~int | ~float64
}

// Type with constraints
type myNumbersConstraints interface {
	constraints.Integer | constraints.Float
}

// Used here in square brackets
func addS[T myNumbers](a, b T) T {
	return a + b
}

func addCS[T myNumbersConstraints](a, b T) T {
	return a + b
}

type myAlias int

func RollTypeConstraint() {
	var n myAlias = 42

	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))

	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.2, 2.2))

	fmt.Println(addS(1, 2))
	fmt.Println(addS(1.2, 2.2))

	fmt.Println(addS(n, 2))
	fmt.Println(addS(1.2, 2.2))

	fmt.Println(addCS(n, 2))
	fmt.Println(addCS(1.2, 2.2))
}
