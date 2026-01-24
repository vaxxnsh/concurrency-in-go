package syncpackage

import (
	"fmt"
	"sync"
)

/*
	Mutex and RW Mutex -
	A Mutex provides a concurrent-safe way to express exclusive access to these shared resources. To borrow
	a Goism, whereas channels share memory by communicating, a Mutex shares memory
	by creating a convention developers must follow to synchronize access to the
	memory.
*/

func MutexExample() {
	var count int
	var lock sync.Mutex
	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}
	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
		for i := 0; i <= 5; i++ {
			arithmetic.Add(1)
			go func() {
				defer arithmetic.Done()
				decrement()
			}()
		}
		arithmetic.Wait()
		fmt.Println("Arithmetic complete.")
	}
}
