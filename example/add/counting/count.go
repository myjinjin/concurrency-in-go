package counting

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func GenerateNumbers(max int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = r.Int()
	}
	return numbers
}

// Add - 숫자를 더하는 순차적인 코드
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

// AddConcurrent - 숫자를 더하는 동시성 코드
func AddConcurrent(numbers []int) int64 {
	// 시스템의 모든 코어 활용
	numOfCores := runtime.NumCPU()
	runtime.GOMAXPROCS(numOfCores)

	var sum int64
	max := len(numbers)

	sizeOfParts := max / numOfCores

	var wg sync.WaitGroup

	for i := 0; i < numOfCores; i++ {
		// Divide the input into parts
		start := i * sizeOfParts
		end := start + sizeOfParts
		part := numbers[start:end]

		// Run computation for each part in separate goroutine.
		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()

			var partSum int64

			// Calculate sum for each part
			for _, n := range nums {
				partSum += int64(n)
			}

			// Add sum of each part to cumulative sum
			atomic.AddInt64(&sum, partSum)
		}(part)
	}
	wg.Wait()
	return sum
}
