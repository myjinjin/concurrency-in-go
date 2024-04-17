package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)

	// 고루틴으로 부터 계산된 값을 얻어오기
	r := <-ch
	fmt.Printf("computed value: %v\n", r)
}
