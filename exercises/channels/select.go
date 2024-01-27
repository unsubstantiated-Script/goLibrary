package channels

import "fmt"

func RollChannelSelect() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go sendz(even, odd, quit)

	receivez(even, odd, quit)

	fmt.Println("about to exit")
}

func sendz(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func receivez(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("from the even channel:", v)
		case v := <-odd:
			fmt.Println("from the odd channel:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("from comma ok:", i, ok)
				return
			} else {
				fmt.Println("from comma ok", i)
			}

		}
	}
}
