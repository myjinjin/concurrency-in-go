package main

import "fmt"

func main() {
	// create channel owner goroutine which return channel and
	// writes data into channel and
	// closes the channel when done.

	owner := func() <-chan int {
		ch := make(chan int) // 채널 생성
		go func() {
			defer close(ch) // 채널 닫기
			for i := 0; i < 5; i++ {
				ch <- i // 채널에 쓰기
			}
		}()
		return ch // 채널 리턴
	}

	consumer := func(ch <-chan int) {
		// read values from channel
		for v := range ch {
			fmt.Printf("Received: %d\n", v)
		}
		fmt.Println("Done receiving!")
	}

	ch := owner()
	consumer(ch)
}

/*
 go run .
Received: 0
Received: 1
Received: 2
Received: 3
Received: 4
Done receiving!
*/
