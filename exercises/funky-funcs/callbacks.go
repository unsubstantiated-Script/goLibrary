package funky_funcs

import "fmt"

func RollCallbacks() {
	x := doMath(1, 2, add)
	fmt.Println(x)
	y := doMath(1, 2, subtract)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}
