package array_slice

import "fmt"

func RollIceCream() {
	iceCreams := [...]string{"Almond Biscotti Café",
		"Banana Pudding ",
		"Balsamic Strawberry (GF)",
		"Bittersweet Chocolate (GF)",
		"Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)",
		"Browned Butter Cookie Dough",
		"Chocolate Covered Black Cherry (GF)",
		"Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)",
		"Coffee (GF)",
		"Cookies & Cream",
		"Fresh Basil (GF)"}

	fmt.Printf("%#v \t %T\n", iceCreams, iceCreams)
	fmt.Println(len(iceCreams))
}
