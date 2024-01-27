package pointers

import "fmt"

// Passing in the pointer to the value
func intDelta(n *int) {
	*n = 43
}

func sliceDelta(ii []int) {
	ii[0] = 99
}

func mapDelta(md map[string]int, s string) {
	md[s] = 65
}

func RollMoarPointers() {
	a := 42
	fmt.Println(a)
	//Gotta pass in the address as the reciever is a pointer
	intDelta(&a)
	fmt.Println(a)

	//This works because a slice is technically a reference pointer to an array.
	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)

	sliceDelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["James"] = 42
	fmt.Println(m["James"])

	mapDelta(m, "James")
	fmt.Println(m["James"])
}
