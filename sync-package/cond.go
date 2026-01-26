package syncpackage

import (
	"fmt"
	"sync"
	"time"
)

/*
	Cond -
	a rendezvous point for goroutines waiting for or announcing the occurrence
	of an event.
*/

func CondExample() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]any, 0, 10)
	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal()
	}
	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
