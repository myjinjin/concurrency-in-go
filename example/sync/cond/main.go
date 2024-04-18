package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup

	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		// suspend goroutine until sharedRsc is populated.
		cond.L.Lock()
		for len(sharedRsc) == 0 {
			cond.Wait()
		}
		cond.L.Unlock()
		fmt.Println(sharedRsc["rsc1"])
	}()

	cond.L.Lock()
	// writes changes to sharedRsc
	sharedRsc["rsc1"] = "foo"
	cond.Signal()
	cond.L.Unlock()
	wg.Wait()
}

/*
 go run .
foo
*/
