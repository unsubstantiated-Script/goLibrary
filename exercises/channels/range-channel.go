package channels

import "fmt"

func RollChannelRange() {
	c := make(chan int)
	//send
	go sendItz(c)

	//receive via range
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("Bounce")
}

// send
func sendItz(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	//Range keeps looping over a channel till its done. Gotta close at the end.
	close(c)
}

//// receive
//func getItz(c <-chan int) {
//	fmt.Println(<-c)
//}
