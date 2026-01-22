package concurrencybuildingblocks

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	Goroutines-
	a goroutine is a function that is running concurrently (remember:
	not necessarily in parallel!) alongside other code. You can start one simply by placing
	the go keyword before a function
*/

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello")
}

func RunConcurrently(wg *sync.WaitGroup) {
	wg.Add(1)
	go sayHello(wg)
}

func MemConsumedInGoroutines() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)
}
