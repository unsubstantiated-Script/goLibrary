package flow

import (
	"fmt"
	"math/rand"
)

func RollStatementIdiom() {

	//Can declare values on the fly inside of if statements
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is Greater than or equal to x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is Less than than x which is %v\n", z, x)
	}

	//The "comma ok" idiom allows us to check values in locations.
	//if seconds, ok := timeZone[tz]; ok {
	//	return "jazz"
	//}
}
