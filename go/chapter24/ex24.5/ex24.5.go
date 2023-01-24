// 일부 고루틴이 데드락이 걸리면서 프로그램이 비정상적인 동작을 한다
// 고루틴이 자원을 서로 다른 순서로 가져가게 해서 생기는 현상이다
// 개발 중에는 데드락이 잘 발생하지 않다가 배포 시에 발생하기 시작함
// 뮤텍스는 가장 단순한 솔루션이지만 성능 문제, 데드락 문제를 고려하여 매우 조심히 사용해야 한다
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}

	// 서로 순서가 다른 경우 데드락 발생
	go diningProblem("A", fork, spoon, "포크", "수저")
	go diningProblem("B", spoon, fork, "수저", "포크")

	//go diningProblem("A", spoon, fork, "수저", "포크")
	//go diningProblem("B", spoon, fork, "수저", "포크")

	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다\n", name)
		// 1000*time.Millisecond 값(1초) 사이의 랜덤한 값을 기다림
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()

}

/* 고루틴 수저, 포크 인자값이 서로 다른 경우 출력 결과
B 밥을 먹으려 합니다
B 수저 획득
B 포크 획득
B 밥을 먹습니다
A 밥을 먹으려 합니다
A 포크 획득
A 수저 획득
A 밥을 먹습니다
B 밥을 먹으려 합니다
A 밥을 먹으려 합니다
A 포크 획득
B 수저 획득
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc000008090?)
        C:/Program Files/Go/src/runtime/sema.go:62 +0x25
sync.(*WaitGroup).Wait(0xf156a9?)
        C:/Program Files/Go/src/sync/waitgroup.go:139 +0x52
main.main()
        C:/Users/박정우/Documents/workspace/github/study/go/chapter24/ex24.5/ex24.5.go:28 +0x171

goroutine 6 [semacquire]:
sync.runtime_SemacquireMutex(0x10?, 0x60?, 0x2?)
        C:/Program Files/Go/src/runtime/sema.go:77 +0x25
sync.(*Mutex).lockSlow(0xc0000160a0)
        C:/Program Files/Go/src/sync/mutex.go:171 +0x165
sync.(*Mutex).Lock(...)
        C:/Program Files/Go/src/sync/mutex.go:90
main.diningProblem({0xfbb846, 0x1}, 0xc000016088, 0xc0000160a0, {0xfbbe3d, 0x6}, {0xfbbe37, 0x6})
        C:/Users/박정우/Documents/workspace/github/study/go/chapter24/ex24.5/ex24.5.go:36 +0x1f1
created by main.main
        C:/Users/박정우/Documents/workspace/github/study/go/chapter24/ex24.5/ex24.5.go:22 +0x10f

goroutine 7 [semacquire]:
sync.runtime_SemacquireMutex(0x10?, 0x60?, 0x2?)
        C:/Program Files/Go/src/runtime/sema.go:77 +0x25
sync.(*Mutex).lockSlow(0xc000016088)
        C:/Program Files/Go/src/sync/mutex.go:171 +0x165
sync.(*Mutex).Lock(...)
        C:/Program Files/Go/src/sync/mutex.go:90
main.diningProblem({0xfbb847, 0x1}, 0xc0000160a0, 0xc000016088, {0xfbbe37, 0x6}, {0xfbbe3d, 0x6})
        C:/Users/박정우/Documents/workspace/github/study/go/chapter24/ex24.5/ex24.5.go:36 +0x1f1
created by main.main
        C:/Users/박정우/Documents/workspace/github/study/go/chapter24/ex24.5/ex24.5.go:23 +0x165
*/

/* 고루틴 수저, 포크 인자 값이 서로 같은 경우 출력 결과
B 밥을 먹으려 합니다
B 수저 획득
B 포크 획득
B 밥을 먹습니다
A 밥을 먹으려 합니다
B 밥을 먹으려 합니다
A 수저 획득
A 포크 획득
A 밥을 먹습니다
A 밥을 먹으려 합니다
B 수저 획득
B 포크 획득
B 밥을 먹습니다
B 밥을 먹으려 합니다
A 수저 획득
A 포크 획득
A 밥을 먹습니다
...
*/

/*
* 그럼 고루틴 호출 함수의 인자 순서만 맞춰주면 데드락 문제가 해결되나?
	* 아니다
	* 고루틴 실행이 if문과 같은 조건문에서의 실행 분기에 따라 다르게 실행되는 경우, 서로 실행 순서가 엇갈릴 수 있다: 데드락 발생
*/
