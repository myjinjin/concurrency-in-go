# Mutex

## When to use channels and when to use mutex

채널과 뮤텍스는 고루틴 간 데이터 공유와 동기화를 위해 사용된다. 채널을 사용하여 고루틴 간에 데이터를 전달하고 작업 단위를 분배하고 비동기 결과를 전달하며, 뮤텍스를 사용하여 캐시, 레지스트리 및 상태를 동시 액세스로부터 보호한다.

각각의 사용 사례와 차이점은 다음과 같다.

- 채널
  - 데이터 복사본을 전달하는 데 사용된다. 여러 고루틴 사이에서 데이터를 안전하게 전달하고 싶을 때 유용하다. 이는 채널이 내부적으로 동기화되어 있어서 동시에 여러 고루틴이 채널에 동시에 접근할 수 없기 때문이다.
  - 작업 단위를 분산시키는 데 사용된다. 예를 들어, 여러 고루틴에게 작업을 할당하고 완료된 결과를 수집할 때 사용된다.
  - 비동기적인 결과를 통신하는 데 사용된다. 채널을 사용하여 고루틴이 완료된 작업의 결과를 다른 고루틴에게 전달할 수 있다.
- 뮤텍스
  - 캐시나 상태와 같은 공유 데이터에 대한 동시 액세스를 관리하는 데 사용된다. 여러 고루틴이 동일한 데이터에 동시에 액세스할 때 데이터 일관성을 보장하기 위해 뮤텍스를 사용한다.
  - 뮤텍스틑 크리티컬 섹션을 보호하는 데 사용된다. 크리티컬 섹션은 여러 고루틴이 동시에 실행할 수 없는 코드 영역을 말한다. 즉, 한 번에 하나의 고루틴만이 크리티컬 섹션에 접근할 수 있다.

## 뮤텍스(Mutex)

- 뮤텍스는 공유 자원에 대한 접근을 보호하기 위해 사용된다.
- 뮤텍스는 개발자들이 따라야 할 규칙을 제공하는데, 공유 메모리에 액세스하고 싶을 때, 먼저 잠금(lock)을 획득하고 완료되면 잠금을 해제해야 한다.
- 잠금은 독점이다. 만약 고루틴이 잠금을 획득했다면, 다른 고루틴은 잠금을 사용할 수 있을 때까지 블락된다.
- 잠금(lock)과 잠금 해제(unlock) 사이의 영역을 크리티컬 섹션(critical section)이라고 한다.
- 잠금 해제는 `defer`와 함께 사용하여, 함수가 끝날 때 잠금 해제를 호출하는 것이 일반적이다.
- 크리티컬 섹션은 고루틴 사이의 보틀넥을 나타낸다.

```go
mu.Lock()
defer mu.Unlock()
balance += amount
```

- `sync.Mutex`: 공유 자원에 독점적인 접근을 제공한다.
- `sync.RWMutex`: 여러 고루틴이 동시에 읽기를 수행할 수 있도록 허용하지만, 쓰기 작업이 이루어지는 동안에는 동시 읽기가 차단된다.

# Atomic

## sync/atomic

- atomic은 메모리에서 로우 레벨 원자 연산을 수행하기 위해 사용된다. 다른 동기화 유틸리티에서도 사용된다.
- 잠금이 없는 작업이다.
- 카운터 원자적 연산에 사용된다.

```go
atomic.AddUint64(&ops, 1)
vaule := atomic.LoadUint64(&ops) // concurrent safe
```

# sync.Cond
 
- 조건 변수는 동기화 메커니즘 중 하나이다.
- 조건 변수는 특정 조건을 기다리는 컨테이너이다.

## 어떻게 하면 고루틴이 이벤트가 발생하거나 어떤 조건이 발생할 때까지 기다리게 할 수 있을까? 

e.g.) 조건이 될 때까지 반복문 안에서 대기한다.

  ```go
  var sharedRsc = make(map[string]string) // 고루틴 간 공유될 맵
  go func() {
    defer wg.Done()
    mu.Lock() // 잠금 획득
    for len(sharedRsc) == 0 { // 공유맵이 비어 있을 때 반복
      mu.Unlock() // 잠금 해제
      time.Sleep(100 * time.Millisecond)
      mu.Lock() // 잠금 획득
    }

    // Do processing...
    fmt.Println(sharedRsc["rsc"])
    mu.Unlock() // 잠금 해제
  }()
  ```

  - 여기서 추가로 필요한 것:
    - 대기 중에 고루틴을 정지할 수 있는 방법이 필요하다.
    - 정지된 고루틴에게 특정한 이벤트가 일어났다는 것을 알리는 어떤 방법이 필요하다.
  - 채널을 통해서 해결할 수는 없나?
    - 채널을 사용하여 수신시 고루틴을 블락할 수 있다.
    - 이벤트 발생을 나타내기 위해 송신 고루틴을 사용할 수 있다.
    - 하지만 여러 조건으로 대기 중인 고루틴이 여러 개 있다면 어떨까?

## sync.Cond

```go
var c *sync.Cond // 조건 변수
```

- 조건 변수는 `sync.Cond` 타입이다.
- 생성자 메서드인 `sync.NewCond()`를 사용하여 조건 변수를 생성하고, 입력으로 `sync.Locker` 인터페이스(일반적으로 sync.Mutex)를 사용한다. 이는 조건 변수가 동시성 안전한 방식으로 고루틴 간의 조정을 용이하게 할 수 있게 해준다.

```go
m := sync.Mutex{}
c := sync.NewCond(&m)
```

- `sync.Cond` 패키지는 3개의 메서드를 포함한다.
  - `c.Wait()`
  - `c.Signal()` 
  - `c.Broadcast()`

## c.Wait()

```go
c.L.Lock()
for !condition {
  c.Wait()
}
// make use of condition ...
c.L.Unlock()
```

- 호출 스레드의 실행을 중지한다.
- 고루틴을 정지하기 전에 자동으로 잠금(`c.L`)을 해제한다.
- Broadcast 또는 Signal이 깨우지 않으면 Wait은 리턴하지 않는다.
- 깨면 다시 잠금(`c.L`)을 획득한다.
- Wait이 처음 재개될 때, `c.L`이 잠겨있지 않기 때문에 호출자는 일반적으로 Wait이 리턴할 때 조건이 true라고 가정할 수 없다. 대신 호출자는 반복문에서 대기해야한다.

## c.Signal()

```go
func (c *Cond) Signal()
```


```go
// G2: shareRsc가 채워질 때까지 기다려야 하는 고루틴
mu := sync.Mutex{}
c := sync.NewCond(&mu)

var shareRsc = make(map[string]string)

go func() {
  defer wg.Done()
  c.L.Lock()
  for len(sharedRsc) == 0 {
    c.Wait() /////////// WAIT!
  }
  // Do processing..
  fmt.Println(sharedRsc["rsc"])
  c.L.Unlock()
}()
```

```go
// G1
go func() {
  defer wg.Done()
  c.L.Lock()
  sharedRsc["rsc"] = "foo"
  c.Signal() // SIGNAL !!!!!
  c.L.Unlock()
}()
```

- Signal은 조건이 되면 대기하고 있던 고루틴 하나를 깨운다.
- Signal은 가장 오래 기다린 고루틴을 찾아내서 그 고루틴에게 알린다.
- 이 호출 중에 호출자가 잠금(`c.L`)을 유지하는 것은 허용되지만 필수는 아니다.

## c.Broadcast()

```go
func (c *Cond) Broadcast()
```

- Broadcast는 조건이 되면 대기하고 있던 모든 고루틴을 께워준다.
- 이 호출 중에 호출자가 잠금(`c.L`)을 유지하는 것은 허용되지만 필수는 아니다.

```go
var sharedRcs = make(map[string]string)

// G2
go func() {
  defer wg.Done()
  c.L.Lock()
  for len(sharedRsc) < 1 {
    c.Wait()  ////// WAIT
  }
  // Do processing
  fmt.Println(sharedRsc["rsc1"])
  c.L.Unlock()
}()

// G3
go func() {
  defer wg.Done()
  c.L.Lock()
  for len(sharedRsc) < 2 {
    c.Wait()  ////// WAIT
  }
  // Do processing
  fmt.Println(sharedRsc["rsc2"])
  c.L.Unlock()
}()

// G1
go func() {
  defer wg.Done()
  c.L.Lock()
  sharedRsc["rsc1"] = "foo"
  sharedRsc["rsc2"] = "bar"
  c.Broadcast()  ////// Broadcast
  c.L.Unlock()
}()
```

## 결론

- 조건 변수는 고루틴 실행의 동기화를 위해 사용된다.
- Wait은 고루틴 실행을 일시정지한다.
- Signal은 조건 c를 기다리는 고루틴 하나를 깨운다.
- Broadcast는 조건 c를 기다리는 모든 고루틴을 깨운다.

# sync.Once

```go
once.Do(funcValue)
```

- `sync.Once`는 한 번의 초기화 함수를 실행하는 데 사용된다.
- `Do` 메서드는 초기화 함수를 인수로 받아들인다.
- `sync.Once`는 다른 고루틴에서 호출된 경우에도 `Do`에 전달된 함수를 한 번만 호출하도록 보장한다.
- 싱글톤 객체를 만들거나, 여러 고루틴에 의존하지만 초기화 함수를 한 번만 호출하기를 원할 때 매우 유용하다.

# sync.Pool

- 풀(Pool)은 일반적으로 데이터베이스 연결, 네트워크 연결, 메모리와 같은 **값비싼 리소스 생성을 제한**하는 데 사용된다.
- 리소스 인스턴스를 고정된 수로 풀을 유지하고, 필요할 때 새 인스턴스를 생성하는 것이 아니라, **풀의 리소스를 재사용**할 것이다.

```go
b := bufPool.Get().(*bytes.Buffer)
// ....
bufPool.Put(b)
```

- 호출자는 리소스에 대한 접근을 원할 때마다 Get() 메서드를 호출할 것이다. 이 메서드는 우선 풀에 사용 가능한 인스턴스가 있는지 확인한다.
- 있으면 해당 인스턴스를 호출자에게 반환한다.
- 없으면 새 인스턴스가 생성되고 해당 인스턴스를 호출자에게 반환한다.
- 사용이 끝나면 호출자는 Put() 메서드를 호출한다. 이 메서드는 인스턴스를 풀에 다시 배치하고, 다른 프로세스에서 재사용할 수 있다.