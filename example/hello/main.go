package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("goroutine-1")

	// goroutine with anonymous function
	go func() {
		fun("goroutine-2")
	}()

	// goroutine with function value call
	fv := fun
	go fv("goroutine-3")

	// wait for goroutines to end
	fmt.Println("wait for goroutines...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done..")
}

/*

 go run .
direct call
direct call
direct call
wait for goroutines...
goroutine-3
goroutine-1
goroutine-2
goroutine-2
goroutine-3
goroutine-1
goroutine-1
goroutine-2
goroutine-3
done..
*/

/*
3개의 고루틴이 동시에 실행되고 있다!
*/
