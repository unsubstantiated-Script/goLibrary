package flow

import "fmt"

func RollSwitchStatements() {

	switch {
	case x < 42:
		fmt.Println("x is less than 42")
	case x == 42:
		fmt.Println("x is equal to 42")
	case x > 42:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println("This is default case")
	}

	//Fallthrough allows switch statement to pass through items to the next condition.
	switch x {
	case 42:
		fmt.Println("x is less than 42")
		fallthrough
	case 43:
		fmt.Println("x is equal to 42")
	case 44:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println("This is default case from fallthrough")
	}

}
