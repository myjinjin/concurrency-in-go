package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fmt.Println(idx)
		}(i)
	}
	wg.Wait()
}

/*
 go run .
3
1
2
*/
