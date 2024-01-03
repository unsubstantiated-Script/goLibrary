package mappage

import (
	"fmt"
	"strconv"
)

func RollMappage() {
	am := map[string]int{
		"Bilbo":     56,
		"Gandalf":   102,
		"Frodizzle": 35,
	}

	fmt.Println("Bilbo was " + strconv.Itoa(am["Bilbo"]) + " when he wrote 'The Hobbit'")

	an := make(map[string]int)
	an["Kakky"] = 32
	an["Shmelp"] = 555
	an["Zips"] = 122
	fmt.Println(an)

	for k, v := range an {
		fmt.Println(k, v)
	}

	delete(an, "Kakky")

	for k, v := range an {
		fmt.Println(k, v)
	}

	if v, ok := an["Dagny"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("key no bueno")
	}
}
