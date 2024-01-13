package funky_funcs

import "fmt"

func RollDefer() {
	//This defer will delay the execution of the function till the end of the other functions. Use case would be closing a file after opening it and making sure that close happens at the end even though the deferred callback was written in a more convenient and readable place like near the open callback.
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
