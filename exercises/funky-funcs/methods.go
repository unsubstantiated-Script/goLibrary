package funky_funcs

import "fmt"

type person struct {
	first string
}

// Here is how the method plays out and the receiver binds it to the struct above.
func (p person) speak() {
	fmt.Println("I am", p.first)
}

func RollMethods() {
	p1 := person{
		first: "James",
	}

	p2 := person{
		first: "Jimbo",
	}

	p1.speak()
	p2.speak()
}
