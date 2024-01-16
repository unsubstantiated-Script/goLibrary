package funky_funcs

import "fmt"

func RollFuncExpressions() {

	x := func(s string) {
		fmt.Println("Gonna yeet this expression", s)
	}

	x("derp")

	y := barz()

	fmt.Println(y())
}

func barz() func() int {
	return func() int {
		return 43
	}
}
