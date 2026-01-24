package syncpackage

import (
	"fmt"
	"sync"
	"time"
)

/*
	Wait Groups -
	WaitGroup is a great way to wait for a set of concurrent operations to complete when
	you either don’t care about the result of the concurrent operation, or you have other
	means of collecting their results.
*/

func WaitGroupExample() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1 * time.Second)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2 * time.Second)
	}()
	wg.Wait()
	fmt.Println("All goroutines complete.")
}
