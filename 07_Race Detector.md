# Go 경쟁 상태 탐지 도구(Go Race Detector)

Go언어의 동시성 메커니즘을 통해 깔끔한 동시성 코드를 비교적 쉽게 작성할 수 있다. 그러나 우리가 실수하는 것을 방지할 수는 없다. 우리가 작성한 코드는 모든 경쟁 조건에 대해 테스트되어야 한다.

Go언어는 코드 내의 경쟁 상태를 감지하기 위한 도구를 제공한다. 이 경쟁 상태 탐지 도구는 다른 Go 도구들과 통합되어 있다.

`go test`, `go run`, `go build`, `go install`에 `-race` 옵션을 사용할 수 있다.

```bash
$ go test -race mypkg # test the package
$ go run -race mysrc.go # compile and run the program
$ go build -race mycmd # build the command
$ go install -race mypkg # install the package
```

- 레이스(race) 도구를 사용하려면 바이너리 빌드가 레이스 활성화 상태여야 한다.
- 바이너리 안에 컴파일러는 모든 메모리 액세스를 기록하기 위한 코드를 삽입한다.
- 런타임 라이브러리는 공유 메모리에 대한 비동기식 액세스를 감시한다. 경쟁적인 동작이 감지되면 경고를 출력한다.
- 레이스 활성화한 바이너리는 10배 더 느리고 10배 더 많은 메모리를 소비한다. 따라서 프로덕션에서 사용할 수는 없다.
- 통합 테스트와 부하 테스트 시, 레이스 활성화된 바이너리로 테스트하면 좋다.

```go
func main() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
```

```bash
$ go build -race main.go
$ ./main
235.697375ms
==================
WARNING: DATA RACE
Read at 0x00c000180020 by goroutine 7:
  main.main.func1()
      /Users/yjlee/myjinjin/concurrency-in-go/example/race/main.go:17 +0xd3

Previous write at 0x00c000180020 by main goroutine:
  main.main()
      /Users/yjlee/myjinjin/concurrency-in-go/example/race/main.go:15 +0x164

Goroutine 7 (running) created at:
  time.goFunc()
      /usr/local/go/src/time/sleep.go:176 +0x47
==================
326.461508ms
964.809311ms
1.670274987s
2.425588476s
2.460326972s
2.582599772s
3.1601991s
3.264886834s
3.468038948s
3.561964987s
4.557566812s
Found 1 data race(s)
```