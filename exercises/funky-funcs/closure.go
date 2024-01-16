package funky_funcs

import "fmt"

func RollClosure() {
	f := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

// Closure: Outer func closing in over an inner func.
func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
