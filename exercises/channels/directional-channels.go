package channels

import "fmt"

func RollDirectionalChannels() {
	c := make(chan int, 2)
	//Recieve only
	//c := make(<-chan int, 2)
	//Send only
	//c := make(chan<- int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("========")
	fmt.Printf("%T\n", c)
}
