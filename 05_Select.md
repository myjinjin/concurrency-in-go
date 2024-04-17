# Select

```go
select {
case <-ch1:
    // block of statements
case <-ch2:
    // block of statements
case ch3 <- struct{}{}:
    // block of statements
}
```

- `select` 구문은 `switch`와 비슷하다.
- 각 case문은 일부 채널에 대한 송신/수신을 지정한다.
- 각 case문은 순차적으로 평가되지는 않지만 모든 채널 작업을 동시에 고려하여 준비된 것이 있는지확인한다.
- 각 case문은 select되는 확률이 동일하다.
- select는 case가 준비될 때까지 기다린다. 채널이 하나도 준비되지 않은 경우 전체 select문이 블락되고, 일부 case가 통신 준비가 될 때까지 기다린다.
- 한 채널이 준비되면 연산이 진행된다.
- 여러 채널이 준비되면 랜덤하게 골라진다.
- select는 타임아웃과 논블로킹 통신 구현에 효율적이다.

# 채널 타임아웃 대기 

```go
select {
case v := <-ch:
    fmt.Println(v)
case <-time.After(3*time.Second):
    fmt.Println("timeout")
}
```

- select는 채널에 이벤트가 있거나 타임아웃이 도달할 때까지 대기한다.
- time.After() 함수는 입력으로 time.Duration이 전달되고, 채널이 반환된다. 이는 백그라운드에서 고루틴을 시작하고, 지정된 시간이 지난 후에 채널에 값을 보낸다.
    ```go
    func After(d Duration) <-chan Time
    ```

# 논블로킹 통신

```go
select {
case m := <-ch:
    fmt.Println("received message", m)
default:
    fmt.Println("no message received")
}
```

- 채널은 블로킹이지만, `select` 연산의 `default case`를 통해 논블로킹을 얻을 수 있다!
- 준비된 채널 작업이 없는 경우, default case가 실행되고 select는 채널을 기다리지 않는다. 작업이 준비되었는지 체크한다. 그렇지 않은 경우 default case가 실행된다.
- 일부 고루틴이 이미 채널 ch에 값을 보낸 경우 값을 읽는다.

# Empty Select

- 비어있는 select문은 영원히 블락된다. 
    ```go
    select {} // blocks forever
    ```
- nil 채널에 select문은 영원히 블락된다.
    ```go
    var ch chan string
    select {
    case v := <-ch: // blocks forever (ch는 nil)
    }
    ```