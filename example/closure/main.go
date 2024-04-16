package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incr := func(wg *sync.WaitGroup) {
		var i int // function의 로컬 변수
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++ // 고루틴 안에서 function의 로컬 변수에 접근
			fmt.Printf("value of i: %v\n", i)
		}()
		fmt.Println("return from function")
	}

	incr(&wg)
	wg.Wait()
	fmt.Println("done..")
}

/*

 go run .
return from function
value of i: 1
done..
*/

// function이 먼저 끝났지만, 고루틴은 여전히 function의 로컬 변수에 접근할 수 있다!
// 보통은 function이 리턴되면, 로컬 변수는 스코프를 벗어난다.
// 그러나 여기서 런타임은 로컬 변수 i에 대한 참조가 여전히 고루틴에 의해 유지되고 있다!
// 영리한 런타임...!!
// 해당 값을 캡쳐해서 스택에서 힙으로 이동시킨다.
// 인클로저 함수가 리턴된 후에도 고루틴이 변수에 대한 접근을 계속할 수 있도록 한다!
