package pointers

import "fmt"

func RollPointerBasics() {
	x := 42
	//Da value
	fmt.Println(x)
	//Da address
	fmt.Println(&x)

	//fmt.Printf("%v\t%T\n", &x, &x)

	y := &x
	fmt.Printf("%v\t%T\n", &y, &y)
	fmt.Println(*y)
}
