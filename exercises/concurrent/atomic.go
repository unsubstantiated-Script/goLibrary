package concurrent

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func RollAtomic() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("CPUs:", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			//Write to counter the Atomic way
			atomic.AddInt64(&counter, 1)
			time.Sleep(time.Second)
			//Read to counter the Atomic way
			fmt.Println("Counter \t", atomic.LoadInt64(&counter))

			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
