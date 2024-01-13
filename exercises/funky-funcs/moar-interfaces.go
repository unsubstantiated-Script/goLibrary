package funky_funcs

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

// Method bound to book is invoking the Stringer library interface
func (b book) String() string {
	return fmt.Sprint("The title of the book is...", b.title)
}

type count int

// Method bound to count is invoking the Stringer library interface
func (c count) String() string {
	//Converting type count to int to string to print it.
	return fmt.Sprint("The count is...", strconv.Itoa(int(c)))
}

func RollInterfaces() {
	//Invoking the structs
	b := book{
		title: "Sashay forth wise man",
	}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)

	log.Println(b)
	log.Println(n)
}
