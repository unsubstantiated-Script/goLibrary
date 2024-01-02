package array_slice

import "fmt"

func RollSlicercises() {
	//arrint := [5]int{}
	//
	//for i := 0; i < 5; i++ {
	//	arrint[i] = i
	//}
	//
	//for i, v := range arrint {
	//	fmt.Printf("%v - %v\n", i, v)
	//}

	//xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//
	//for i, v := range xi {
	//	fmt.Printf("%v - %v\n", i, v)
	//}
	//
	//x1 := xi[:5]
	//fmt.Println(x1)
	//
	//x2 := xi[5:]
	//fmt.Println(x2)
	//
	//x3 := xi[2:7]
	//fmt.Println(x3)
	//
	//x4 := xi[1:6]
	//fmt.Println(x4)

	//x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//x = append(x, 52)
	//fmt.Println(x)
	//x = append(x, 53, 54, 55)
	//fmt.Println(x)
	//
	//y := []int{56, 57, 58, 59, 60}
	//x = append(x, y...)
	//fmt.Println(x)

	//x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//x = append(x[:3], x[6:]...)
	//fmt.Println(x)

	//states := make([]string, 0, 50)
	//
	//states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	//
	//fmt.Println(len(states))
	//fmt.Println(cap(states))
	//
	//for i := 0; i < len(states); i++ {
	//	fmt.Printf("%v - %v\n", i, states[i])
	//}

	jb := []string{
		"James",
		"Bond",
		"Sprite",
		"Chocolate",
	}

	jm := []string{
		"Jenny",
		"Moneypenny",
		"Coke",
		"Rocky Road",
	}

	spies := [][]string{jb, jm}

	for i, v := range spies {
		fmt.Println(i, v)
		for j, s := range v {
			fmt.Println(j, s)
		}

	}
}
