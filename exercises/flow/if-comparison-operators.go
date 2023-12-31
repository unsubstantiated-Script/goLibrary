package flow

import "fmt"

func RollComparisonLogic() {

	fmt.Printf("x=%v \n y=%v\n", x, y)

	// Conditionals
	// If statments
	// Switch statements

	if x < 42 {
		fmt.Println("less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("less than the meaning of life")
	} else {
		fmt.Println("moar than the meaning of life")
	}

	if x < 42 {
		fmt.Println("less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equal to the meaning of life")
	} else {
		fmt.Println("moar than the meaning of life")
	}
}
