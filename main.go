package main

import (
	"sync"

	concurrencybuildingblocks "github.com/vaxxnsh/concurrency-in-go/concurrency-building-blocks"
)

func main() {

	var wg sync.WaitGroup

	for range 5 {
		concurrencybuildingblocks.RunConcurrently(&wg)
	}

	wg.Wait()
}
