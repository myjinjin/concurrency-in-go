package counting

import (
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	numbers := GenerateNumbers(1e7)
	for i := 0; i < b.N; i++ {
		Add(numbers)
	}
}

func BenchmarkAddConcurrent(b *testing.B) {
	numbers := GenerateNumbers(1e7)
	for i := 0; i < b.N; i++ {
		AddConcurrent(numbers)
	}
}

/*
 go test ./... -bench=.
?       github.com/myjinijn/concurrency-in-go/example/add      [no test files]
goos: darwin
goarch: amd64
pkg: github.com/myjinijn/concurrency-in-go/example/add/counting
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkAdd-8                       372           3292165 ns/op
BenchmarkAddConcurrent-8             493           2558118 ns/op
PASS
ok      github.com/myjinijn/concurrency-in-go/example/add/counting      4.588s
*/
