# Introduction

## 동시성(Concurrency)

- 동시성은 여러 가지 일이 동시에 무작위 순서로 일어나는 것을 의미한다.
- Go는 동시성에 대한 지원을 기본적으로 제공한다.

## Why we need to think about concurrency?

```go
// Add - 숫자를 더하는 순차적인 코드
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}
```

이 function을 더 빠르게 실행하는 방법이 없을까?

## Computing Environment

- 멀티코어 프로세서 (Add, idle, idle, idle)
- `Add()`는 싱글코어에서 실행된다. 나머지 코어들은 여전히 idle 상태이다.
- 다른 코어도 사용하여 계산을 더 빠르게 하고 싶다.
- 한 가지 방법은, 입력값을 나눠서 다른 코어에서 병렬적으로 각 파트에 `Add()` 함수의 여러 인스턴스를 실행할 수 있다.

```go
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
```

## 동시성(Concurrency) vs 병렬성(Parallelism)

- 동시성은 실행 중인 프로세스들을 결합하는 행위이다.
  - e.g. OS는 마우스, 키보드, 스크린 등 다양한 주변 장치를 제어하지만 이러한 여러 개의 장치를 관리하기 위해서 병렬 작업이 필요한 건 아니다. 하나의 CPU 코어로도 처리가 가능하다.
- 여러 개의 작업을 다루는 것에 관한 개념이다.
- 동시성을 통해 병렬성을 달성할 수 있긴 하지만 동시성의 목적이 병렬성은 아니다.
- 동시성의 목적은 구조화에 있다.
- 동시성은 프로그램을 여러 개의 작은 부분으로 나누어 각각의 부분이 서로 정보를 주고 받으며 독립적으로 실행될 수 있도록 하는 구조화 작업이다.
- 병렬성은 여러 작업을 동시에 실행할 수 있는 능력이다.
  - e.g. 두 개의 코어가 각기 다른 쓰레드/프로세스를 실행한다.

# Processes and Threads

# Why Concurrency is hard