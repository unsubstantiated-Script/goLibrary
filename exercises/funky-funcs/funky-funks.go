package funky_funcs

import "fmt"

func RollFunkyFuncs() {

	//Basic structure
	//func (receiver) identifier(params) (returns) {code}

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//This won't take a slice xi in, it wants a variatic or a group of numbers. So it needs to be unfurled.
	x := sumNums(xi...)

	fmt.Println("The sum is", x)
}

// Variatic ... unlimited or unknown amount in argument. Still have to declare type. Go turns into a slice.
func sumNums(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)
	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}
