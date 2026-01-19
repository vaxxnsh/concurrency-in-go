package main

import (
	"fmt"

	"github.com/vaxxnsh/concurrency-in-go/synchronization"
)

func main() {

	fmt.Println("showing example for uncertain concurrent code -")

	for range 5 {
		synchronization.DeadlockExample()
	}
}
