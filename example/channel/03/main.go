package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 6)

	go func() {
		defer close(ch)

		// send all iterator values on channel without blocking
		for i := 0; i < 6; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}
}

/*
ch := make(chan int) 일 때 출력

 go run .
Sending: 0
Sending: 1
Received: 0
Received: 1
Sending: 2
Sending: 3
Received: 2
Received: 3
Sending: 4
Sending: 5
Received: 4
Received: 5
*/

/*
ch := make(chan int, 6) 일 때 출력
-> 블로킹 없이 6개 바로 전송 가능!
 go run .
Sending: 0
Sending: 1
Sending: 2
Sending: 3
Sending: 4
Sending: 5
Received: 0
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
*/
