package array_slice

import "fmt"

func RollSlices() {
	////Like an array w/o number in square
	//xs := []string{"hello", "toad muffin", "of doom"}
	////Appending to a slice
	//xss := append(xs, "derpleing")
	//fmt.Println(xss)
	//
	////Slicing a slice
	//xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//
	//fmt.Printf("xi - %#v\n", xi)
	//fmt.Println("----------------")
	//
	//// [inclusive:exclusive]
	//fmt.Printf("inclusive:exclusive xi - %#v\n", xi[1:3])
	//
	//// [:exclusive]
	//fmt.Printf("exclusive xi - %#v\n", xi[:7])
	//
	//// [inclusive:]
	//fmt.Printf("inclusive xi - %#v\n", xi[4:])
	//
	////Slicing a slice or deleting items from a slice
	////Second argument is "unfurling the slice"
	//xsss := append(xi[:4], xi[5:]...)
	//fmt.Printf("delete xsss - %#v\n", xsss)
	//
	////Slice make
	//xs4 := make([]int, 0, 10)
	//fmt.Println(xs4)
	//fmt.Println(len(xs4))
	//fmt.Println(cap(xs4))
	//
	//xs4 = append(xs4, 33, 43, 42, 88, 77)
	//fmt.Println(xs4)

	//Multi-dimensional Slice
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

	fmt.Println(jb)
	fmt.Println(jm)

	//Slice of slices. That's enough slices!
	xxs := [][]string{jb, jm}

	fmt.Println(xxs)
}
