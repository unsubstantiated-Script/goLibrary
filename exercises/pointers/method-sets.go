package pointers

import "fmt"

type dog struct {
	first string
}

// Method attached to the struct using Value Semantics
func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

// Method attached to the struct using Pointer Semantics
func (d *dog) run() {
	//Was able to mutate the value here
	d.first = "Cheese"
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}
func RollMethodSets() {

	//These will show that method calls will still where here on method sets.
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()
	youngRun(&d1)

	d2 := &dog{"Deetz"}
	d2.walk()
	d2.run()
	youngRun(d2)

}
