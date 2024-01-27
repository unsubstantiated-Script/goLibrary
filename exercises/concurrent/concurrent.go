package concurrent

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func RollConcurrent() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	//Go Routines //WaitGroup allows us to wait for one routine do do it's thang
	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
