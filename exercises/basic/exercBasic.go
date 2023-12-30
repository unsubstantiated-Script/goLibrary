package basic

import "fmt"

var caliban string

const cheese string = "Cheddar"

func PrintStuffs() {
	crackers := "Ritz"

	caliban = "man"

	fmt.Printf("The %s Caliban ate hisself some %s and %s \n", caliban, crackers, cheese)
}
