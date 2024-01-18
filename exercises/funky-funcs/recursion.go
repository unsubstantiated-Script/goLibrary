package funky_funcs

import "fmt"

func RollRecursion() {
	fmt.Println(4 * 3 * 2 * 1)
	fmt.Println(factorial(4))
	fmt.Println(factLoop(4))
}

func factorial(n int) int {
	fmt.Println("This is n", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Recursion can also be done with loops, but loops tend to be easier to understand. Recursion is probably best suited for grid looping, but still not as efficient as a filter or something.
func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}
