package main

import (
	"fmt"
	"github.com/unsubstantiated-Script/goDawggy"
)

func main() {
	s1 := goDawggy.Bark()
	s2 := goDawggy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)
}
