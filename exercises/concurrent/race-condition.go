package concurrent

import (
	"fmt"
	"runtime"
	"sync"
)

func RollRaceCondition() {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("CPUs:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	//Using a Mutex to lock down code in a GoRoutine you don't want each routine to touch.
	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			//Due to goroutines, this "v" will be different every time...
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
