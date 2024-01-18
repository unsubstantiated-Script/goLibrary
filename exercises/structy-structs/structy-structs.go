package structy_structs

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// Embedding structs vs inheriting
type secretAgent struct {
	person
	ltk bool
}

func RollStruct() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	//Anonymous struct defines and loads.
	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "Bilbo",
		last:  "Baggins",
		age:   38,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Printf("%T %#v\n", sa1, sa1)
	fmt.Printf("%T %#v\n", p2, p2)
	//fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk, sa1.person)
}
