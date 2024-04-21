# Interface values

개념적으로 인터페이스 타입의 값에는 동적 타입과 동적 값이라는 두 가지 구성 요소가 있다.

- 동적 타입(Dynamic type)
- 동적 값(Dynamic value)

```go
var d Dense
```

![06-interface-dynamic-type-value.svg](./images/06-interface-dynamic-type-value.svg)

인터페이스 타입의 변수를 선언하면 동적 타입과 동적 값이 nil로 설정된다.

```go
d = &gold
```

![06-interface-dynamic-type-value2.svg](./images/06-interface-dynamic-type-value2.svg)


변수 gold에 대한 포인터 값을 할당하면 인터페이스 동적 타입이 Metal 포인터 타입인 타입 디스크립터로 설정된다.

```go
d.Density()
```

![06-interface-dynamic-type-value3.svg](./images/06-interface-dynamic-type-value3.svg)

인터페이스를 통해 메서드를 호출하면 동적 디스패치(dynamic dispatch)를 사용한다. 컴파일러는 타입 디스크립터로부터 메서드 주소를 얻기 위해 코드를 생성한다. 따라서 런타임에 해당 주소에서 간접 호출이 이루어진다.

호출에 대한 리시버 인수는 인터페이스 동적 값의 복사본이다.

![06-interface-dynamic-type-value4.svg](./images/06-interface-dynamic-type-value4.svg)

```go
d.Density()
// ->
gold.Density
```

따라서 인터페이스를 통해 메서드를 호출하면 해당 타입의 값에서 메서드를 호출하는 것과 동일해진다.

Gas 타입의 값을 인터페이스에 할당했을 때, 동적 타입이 Gas 타입에 대한 포인터로 설정되고 동적 값이 oxygen 변수의 값을 가리킬 때도 비슷하게 동작한다.

```go
d = &oxygen
```

![06-interface-dynamic-type-value5.svg](./images/06-interface-dynamic-type-value5.svg)


```go
d.Density()
```

![06-interface-dynamic-type-value6.svg](./images/06-interface-dynamic-type-value6.svg)

Density() 메서드를 호출하면 타입 디스크립터의 메서드 주소를 사용하여 동적 값을 리시버 값으로 사용하는 메서드를 간접적으로 호출한다.

# Interface의 목적

## 캡슐화(Encapsulate)

인터페이스를 사용하면 사용자 정의 타입의 메서드 내에 로직을 캡슐화할 수 있다.

```go
func (m *Metal) Density() float64 {
    return m.mass / m.volmue
}

func (g *Gas) Density() float64 {
    var density float64
    density = (g.molecularMass * g.pressure) / (0.0821 * (g.temperature + 273))
    return density
}
```

데이터 타입 Metal과 Gas는 각각 밀도를 계산하는 고유한 공식이 있다. 따라서 해당 로직을 메서드 내에 캡슐화할 수 있다.

## 추상화(Abstraction)

인터페이스는 기본 구체적인 타입의 동작을 보장하는 상위 수준 함수에 대한 추상화를 제공하므로 상위 수준 함수가 특정 구현의 세부 사항에 묶여있지 않으므로 다양한 코드를 작성할 수 있다.

```go
func IsDenser(a, b Dense) bool {
    return a.Density() > b.Density()
}

result := IsDenser(&gold, &silver)

result = IsDenser(&oxygen, &hydrogen)
```

# 암시적 인터페이스(Interface-implicit)

암시적 인터페이스는 좋은 디자인으로 이어진다.

## 인터페이스(Interfaces)

- 인터페이스는 암묵적으로(implicitly) 충족된다!
- Java에서는 명시적으로 인터페이스를 구현하는 것으로 클래스를 선언하지만 Go에서는 그러한 명시적인 구문이 없다.

아래는 Java에서의 인터페이스 구현 예시이다.

```java
class Bicycle implements Vehicle{
```

- 사용자 정의 타입은 인터페이스의 인스턴스로 간주되기 위해 인터페이스에 정의된 메서드를 소유하기만 하면 된다.
- 인터페이스의 정의는 구현과 분리된다. 이는 많은 유연성을 제공한다.
- 구현 전에 인터페이스를 정의해야 함으로써 프로젝트 시작 단계에서 추상화에 얽매일 필요가 없다!
- Go에서는 추상화가 명확해지면 인터페이스를 정의할 수 있다.
- 이 설계를 통해 우리는 기존의 구체적인 타입을 충족하는 새로운 인터페이스를 만들 수 있다. 모든 구현으로 돌아가서 태그를 지정할 필요가 없다. 표준 라이브러리나 서드파티 코드처럼 구현이 통제를 받지 않는 경우 때로는 불가능할 수도 있다.
- 인터페이스 정의와 구체적인 타입 정의는 사전 조정 없이 모든 패키지에 나타날 수 있다. 따라서 다른 패키지를 가져올 필요 없이 원하는 서명으로 인터페이스를 정의할 수 있는 새 패키지를 구현할 때 종속성을 추상화하는 것이 더 쉬워진다.

## 컨벤션(Convention)

- 인터페이스를 단순하고 짧게 유지한다.
- 통일된 방식으로 처리해야 하는 두 개 이상의 구체적인 타입이 있는 경우 인터페이스를 정의한다.
- 더 적고 간단한 방법으로 더 작은 인터페이스를 만든다.
- 인터페이스 설계에 대한 좋은 경험 법칙은 필요한 것만 요청하는 것이다.

# 구체적인 타입과 값 숨기기

- 인터페이스는 구체적인 타입과 값을 숨긴다.
- 인터페이스는 구체적인 타입을 래핑하며, 구체적인 타입이 다른 메서드를 구현하더라도 인터페이스에 정의된 메서드만 표시된다.

# 타입 단언(Type assertion)

- 타입 단언은 인터페이스 값에 적용되는 작업이다.
- 인터페이스 값에서 동적 값을 추출하는 데 사용된다.

```go
v := x.(T)
```

- x: 인터페이스 타입
- T: 추출하려넌 값을 갖는 구체적인 타입

이 작업은 x의 동적 타입이 구체적인 타입 T와 동일한지 확인하고, 동일하다면 동적 값을 반환한다. 실패하면 작업이 패닉 상태가 된다.

```go
v, ok := x.(T)
```

- ok는 boolean 값을 리턴한다. 이는 지정된 타입이 인터페이스 값에 포함되어 있는지 여부를 나타낸다.
- 여기서는 타입 단언이 실패하더라도 패닉이 발생하지 않는다.

```go
func shutdownWrite(conn net.Conn) {
    v, ok := conn.(*net.TCPConn)
    if ok {
        v.CloseWrite() // TCPConn에만 정의되어 있는 메서드
    }
}
```

# 빈 인터페이스(Empty Interface)

Go 코드 많은 곳에서 사용되는 빈 인터페이스를 볼 수 있다.

```go
func fmt.Println(a ...interface{}) (n int, err error)
```

```go
func fmt.Errorf(format string, a ...interface{}) error
```

- 빈 인터페이스 타입은 이를 충족하는 타입을 요구하지 않는다. 구현할 방법이 없다.
- 빈 인터페이스에 어떤 값이든 할당할 수 있다.
- Println 및 Errorf 함수는 빈 인터페이스를 파라미터로 사용하므로 Go의 모든 유형을 입력으로 전달할 수 있다.
- 그리고 내부적으로는 타입 스위치(type switch)이나 타입 단언(type assertion)을 사용하여 타입을 구별한다.

## 타입 스위치(Type Switch)

```go
switch v := value.(type) {
case int:
    fmt.Printf("v is integer with value %d\n", v)
case string:
    fmt.Printf("v is a string, whose length is %d\n", len(v))
default:
    fmt.Println("we dont know what 'v' is!")
}
```

타입 스위치는 인터페이스 변수의 동적 타입을 찾는 데 사용된다. 타입 스위치는 괄호 안에 `type` 키워드를 사용하는 타입 단언 구문을 사용한다.

## 빈 인터페이스 주의점

빈 인터페이스는 함수에 들어오는 데이터에 대한 정보를 제공하지 않으므로 사용할 때 주의해야 한다.

정적으로 유형이 지정된 언어의 이점이 무효화되고 컴파일러는 더 이상 우리가 실수를 하여 잘못된 데이터 유형을 함수에 전달했다는 것을을 알려줄 수 없다.

대부분의 경우 빈 인터페이스를 사용하기 보다는 특정 데이터 타입을 사용하거나 필요한 특정 방법을 사용하여 인터페이스를 만드는 것이 좋다.