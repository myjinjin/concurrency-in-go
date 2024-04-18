package main

import (
	"fmt"
	"math/rand"
	"time"
)

// identify the data race
// fix the issue.

func main() {
	start := time.Now()
	var t *time.Timer
	ch := make(chan bool)
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		ch <- true
	})
	for time.Since(start) < 5*time.Second {
		<-ch
		t.Reset(randomDuration())
	}
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

//----------------------------------------------------
// (main goroutine) -> t <- (time.AfterFunc goroutine)
//----------------------------------------------------
// (working condition)
// main goroutine..
// t = time.AfterFunc()  // returns a timer..

// AfterFunc goroutine
// t.Reset()        // timer reset
//----------------------------------------------------
// (race condition- random duration is very small)
// AfterFunc goroutine
// t.Reset() // t = nil

// main goroutine..
// t = time.AfterFunc()
//----------------------------------------------------

/*
$ go build -race main.go
$ ./main
141.201251ms
249.543782ms
673.489254ms
1.557176956s
2.454573628s
2.577232121s
3.150769991s
3.314358772s
4.053545955s
4.87727441s
5.333257105s
5.735393199s
*/
