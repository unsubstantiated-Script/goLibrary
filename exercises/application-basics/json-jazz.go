package application_basics

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func RollJSONJazz() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	write, err := os.Stdout.Write(b)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(write)

	var jsonBlob = []byte(`[
{"Name": "Calcifer", "Order":"Fire"},
{"Name": "Specks", "Order":"Vision"}
]`)

	type Animal struct {
		Name  string
		Order string
	}

	var animals []Animal
	
	err = json.Unmarshal(jsonBlob, &animals)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v", animals)

}
