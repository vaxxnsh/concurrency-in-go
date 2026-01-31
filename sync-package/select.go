package syncpackage

import (
	"fmt"
	"time"
)

/*
	Select-
	The select statement is the glue that binds channels together; it’s how we’re able to
	compose channels together in a program to form larger abstractions. If channels are
	the glue that binds goroutines together, what does that say about the select statement?
	It is not an overstatement to say that select statements are one of the most
	crucial things in a Go program with concurrency
*/

func SelectExample() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}
