package concurrencybuildingblocks

import (
	"fmt"
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
