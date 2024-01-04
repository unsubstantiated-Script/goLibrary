package structy_structs

import "fmt"

type perp struct {
	first   string
	last    string
	iceFlav []string
}

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	doors  int
	color  string
}

func RollStructercises() {
	//p1 := perp{
	//	first: "JoJo",
	//	last:  "Potatoes",
	//	iceFlav: []string{
	//		"killa nilla",
	//		"strew berry",
	//		"plain jane",
	//	},
	//}
	//
	//p2 := perp{
	//	first: "Skippity",
	//	last:  "Bopps",
	//	iceFlav: []string{
	//		"Chew Kinking",
	//		"schonk monk",
	//		"turdle",
	//	},
	//}

	//fmt.Println(p1)
	//fmt.Println(p2)
	//
	//fmt.Println(p1.iceFlav)
	//fmt.Println(p2.iceFlav)
	//
	//for _, v := range p1.iceFlav {
	//	fmt.Println(p1.first, "favorite is", v)
	//}
	//
	//for _, v := range p2.iceFlav {
	//	fmt.Println(p2.first, "favorite is", v)
	//}

	//per1 := map[string]perp{
	//	p1.last: p1,
	//	p2.last: p2,
	//}
	//
	//per2 := map[string]perp{
	//	p1.last: p1,
	//	p2.last: p2,
	//}
	//
	//for _, v := range per1 {
	//	fmt.Println(v)
	//	for _, v2 := range v.iceFlav {
	//		fmt.Println(v.first, v.last, v2)
	//	}
	//}
	//
	//for _, v := range per2 {
	//	fmt.Println(v)
	//	for _, v2 := range v.iceFlav {
	//		fmt.Println(v.first, v.last, v2)
	//	}
	//}

	//c1 := vehicle{
	//	engine: engine{
	//		electric: true,
	//	},
	//	make:  "Doodge",
	//	model: "Caliper",
	//	doors: 5,
	//	color: "pahnk",
	//}
	//
	//c2 := vehicle{
	//	engine: engine{
	//		electric: false,
	//	},
	//	make:  "Chewvy",
	//	model: "k20",
	//	doors: 2,
	//	color: "tahneal",
	//}
	//
	//fmt.Println(c1)
	//fmt.Println(c2)
	//fmt.Println(c1.engine.electric)
	//fmt.Println(c2.engine.electric)
	//fmt.Println(c1.model)
	//fmt.Println(c2.model)

	p2 := struct {
		first    string
		friends  map[string]int
		favDrank []string
	}{
		first: "Bilbo",
		friends: map[string]int{
			"Zambo":   4,
			"Kingsly": 55,
		},
		favDrank: []string{"Coke", "Dew", "Chaulk"},
	}

	fmt.Println(p2.first, p2.friends, p2.favDrank)

	for k, v := range p2.friends {
		fmt.Println(p2.first, k, "- friends -", v)
	}

	for _, v := range p2.favDrank {
		fmt.Println(p2.first, v)
	}

}
