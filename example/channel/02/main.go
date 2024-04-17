package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			// send iterator over channel
			ch <- i
		}
		close(ch) // close하지 않으면 fatal error: all goroutines are asleep - deadlock! 발생
	}()

	// range over channel to recv values
	for v := range ch {
		fmt.Println(v)
	}
}
