package funky_funcs

import "fmt"

func RollAnonymoose() {
	fooby()

	//Self instantiated anonymous funcs
	func() {
		fmt.Println("Anonymous!")
	}()

	func(s string) {
		fmt.Println("Anonymous name is", s)
	}("Nodd")
}

func fooby() {
	fmt.Println("Foo ran")
}
