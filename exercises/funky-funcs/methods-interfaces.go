package funky_funcs

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

// This interface allows polymorphism. Hey baby, if you've got these methods, then you're my type ;X
type human interface {
	speak()
}

// Here is how the method plays out and the receiver binds it to the struct above.
func (p person) speak() {
	fmt.Println("I am", p.first)
}

// Here is how the method plays out and the receiver binds it to the struct above.
func (sa secretAgent) speak() {
	ltk := determineLTK(sa.ltk)
	fmt.Printf("I am %v a secret agent man. I %v have a license to kill.\n", sa.person.first, ltk)
}

func determineLTK(ltkBool bool) string {
	ltk := ""

	if ltkBool {
		ltk = "do"
	} else {
		ltk = "do not"
	}

	return ltk
}

// interfacing with the interface
func saySomething(h human) {
	h.speak()
}

func RollMethods() {
	p1 := person{
		first: "James",
	}

	p2 := person{
		first: "Jimbo",
	}

	//p1.speak()
	//p2.speak()

	saySomething(p1)
	saySomething(p2)

	sa1 := secretAgent{
		person: person{
			first: "Skippy",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Danger",
		},
		ltk: false,
	}

	//sa1.speak()
	//sa2.speak()

	saySomething(sa1)
	saySomething(sa2)
}
