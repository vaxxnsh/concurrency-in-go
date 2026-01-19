package synchronization

import "fmt"

/*
	Race Condition -
	A race condition occurs when two or more operations must execute in the correct
	order, but the program has not been written so that this order is guaranteed to be
	maintained.
*/

/*
	Atomicity-
	When something is considered atomic, or to have the property of atomicity, this
	means that within the context that it is operating, it is indivisible, or uninterruptible.

	Most statements are not atomic, let alone functions, methods, and programs.
*/

var data int

func BadConcurrentCode() {

	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
}
