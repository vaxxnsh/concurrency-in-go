package synchronization

import (
	"fmt"
	"sync"
)

var memoryAccess sync.Mutex

func MemoryAccessSync() {

	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
	memoryAccess.Unlock()
}
