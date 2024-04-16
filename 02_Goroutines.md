# Goroutines

## Communicating Sequential Processes(CSP)

Go에서 동시성은 Communicating Sequential Processes(CSP)에 대해 Tony Hoare가 작성한 논문이 바탕이 된다. CSP는 매우 간결하다. CSP는 3가지 핵심 아이디어를 기반으로 한다.

- 각 프로세스는 **순차적 실행**(Sequential execution)으로 구성된다.
  - 모든 프로세스에는 로컬 상태가 있으며 프로세스는 해당 로컬 상태에서 작동한다.
- 한 프로세스에서 다른 프로세스로 데이터를 전송해야 하는 경우 **메모리를 공유하지 않고**(No shared memory) 데이터를 통신하여 다른 프로세스로 데이터 사본을 보낸다.
  - 메모리 공유가 없기 때문에 경쟁 상태, 데드락이 없다.
- 각 프로세스가 독립적으로 실행될 수 있기 때문에 쉽게 **확장**(Scale)할 수 있다.
  - 계산에 더 많은 시간이 소요되는 경우 동일한 유형의 프로세스를 더 추가하고 계산을 더 빠르게 실행할 수 있다.

## Go's Concurrency Tool Set

- goroutines
- channels
- select
- sync package

## 고루틴(Goroutines)

- 고루틴은 Go runtime에 의해 관리되는 **사용자 공간 스레드**(user space threads)이다.
- Go runtime은 실행 파일의 일부이다. 애플리케이션의 실행 파일에 내장되어 있다.
- 고루틴은 매우 가볍다.
- 고루틴은 필용에 따라 성장하고 축소할 수 있는 **2KB의 스택**(2KB of stack)으로 시작한다.
- **CPU 오버헤드가 매우 낮다**.(Low CPU overhead)
- 고루틴을 만드는 데 필요한 CPU 명령의 양이 매우 적다.
- 동일한 주소 공간에 **수십만 개의 고루틴**을 만들 수 있다.
- 고루틴 간에 **채널을 사용하여 데이터를 전달**하여, 메모리 공유를 피할 수 있다.
- 스레드 컨텍스트 전환에 비해 고루틴 **컨텍스트 전환 비용이 적다**.
- Go runtime은 데이터를 영속화하고 검색하기 위해 어떤 것을 영속화할지, 어떻게 영속화할지, 그리고 언제 영속화해야 하는지에 대해 보다 선택적으로 작동할 수 있다.
- Go runetime은 OS 스레드를 생성한다.
- 고루틴은 OS 스레드의 컨텍스트에서 실행된다. **** 중요 *****

![01-goroutines](./images/01-goroutines.png)

- 많은 고루틴이 단일 OS 스레드의 컨텍스트에서 실행될 수 있다.
- 운영체제는 OS 스레드를 스케줄링한다.
- Go runtime은 OS 스레드에서 여러 고루틴의 스케줄링을 관리한다.
- 운영체제는 변경된 것이 없으며 여전히 스레드를 예약하고 있다.

![01-goroutines-on-os-thread.png](./images/01-goroutines-on-os-thread.png)

## 결론

### 고루틴(Goroutine)이란

- Go runtime에 의해 관리되는 사용자 공간 스레드이다.

### 고루틴의 장점 (OS 스레드에 비하여)

- 고루틴은 OS 스레드에 비해 매우 경량이다.
- OS 스레드의 스택 사이즈인 8MB에 비해 고루틴의 스택 사이즈는 2KB로 매우 작다.
- 사용자 공간에서 발생하기 때문에 컨텍스트 전환 비용이 매우 적다. 고루틴은 저장할 상태가 매우 적다.
- 동일한 머신에서 수십만 개의 고루틴을 생성할 수 있다.
