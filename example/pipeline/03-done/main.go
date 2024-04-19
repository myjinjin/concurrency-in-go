// generator() -> square() ->
//
//															-> merge -> print
//	            -> square() ->
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func generator(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	done := make(chan struct{})

	in := generator(done, 2, 3) // 첫 번째 파라미터로 done 채널을 쓰는 것이 컨벤션이다!

	c1 := square(done, in)
	c2 := square(done, in)

	out := merge(done, c1, c2)

	// cancel goroutines after receiving one value.
	fmt.Println(<-out)
	close(done) // 고루틴을 취소하기 위해 done 채널을 사용했다
	time.Sleep(2 * time.Millisecond)

	// how to check if the goroutines have terminated
	g := runtime.NumGoroutine()
	fmt.Printf("number of goroutines active = %d\n", g)
}
