package main

import (
	"fmt"
	"time"

	"github.com/myjinijn/concurrency-in-go/example/add/counting"
)

func main() {
	numbers := counting.GenerateNumbers(1e7)

	t := time.Now()
	sum := counting.Add(numbers)
	fmt.Printf("Sequential Add, Sum: %d, Time Taken: %s\n", sum, time.Since(t))

	t = time.Now()
	sum = counting.AddConcurrent(numbers)
	fmt.Printf("Concurrent Add, Sum: %d, Time Taken: %s\n", sum, time.Since(t))
}

/*
 go run ./example/add/.
Sequential Add, Sum: -88775338718211922, Time Taken: 7.345102ms
Concurrent Add, Sum: -88775338718211922, Time Taken: 1.810557ms
*/
