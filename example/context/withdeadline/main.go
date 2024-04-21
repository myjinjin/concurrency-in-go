package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {
	// set deadline for goroutine to return computational result.
	deadline := time.Now().Add(100 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	compute := func(ctx context.Context) <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)
			deadline, ok := ctx.Deadline()
			if ok {
				if deadline.Sub(time.Now().Add(50*time.Millisecond)) < 0 {
					fmt.Println("not sufficient time given... terminating")
					return
				}
			}

			time.Sleep(50 * time.Millisecond)

			select {
			case ch <- data{"123"}: // Report result.
			case <-ctx.Done():
				return
			}
		}()
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	ch := compute(ctx)

	d, ok := <-ch
	if ok {
		fmt.Printf("work is complete %s\n", d)
	}
}
