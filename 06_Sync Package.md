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