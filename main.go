package main

import (
	"sync"

	syncpackage "github.com/vaxxnsh/concurrency-in-go/sync-package"
)

func main() {

	var wg sync.WaitGroup

	for range 5 {
		syncpackage.WaitGroupExample()
	}

	wg.Wait()
}
