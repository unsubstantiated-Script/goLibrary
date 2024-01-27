package channels

import "fmt"

func RollDemChannels() {

	c := make(chan int)
	//send
	go sendIt(c)

	//receive
	getIt(c)

	fmt.Println("Bounce")
}

// send
func sendIt(c chan<- int) {
	c <- 42
}

// receive
func getIt(c <-chan int) {
	fmt.Println(<-c)
}
