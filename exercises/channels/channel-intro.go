package channels

import "fmt"

func RollChannels() {
	c := make(chan int)

	//Buffer channel <- don't use em all that often
	c2 := make(chan int, 1)

	c2 <- 84

	//Putting jazz on the channel

	go func() {
		c <- 42
	}()

	//Taking jazz off the channel. Won't run as it's just opening a channel to do stuffs.
	//Channels block!!
	fmt.Println(<-c)
	fmt.Println(<-c2)
}
