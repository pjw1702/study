// 고루틴의 동시성 문제를 해결할 수 있는 가장 간단한 방법으로서, 뮤텍스를 이용해 락과 언락을 거는 것이 있다
// 프로그램이 정상적으로 돌아가므로 출력 값이 나오지 않는 상태로 무한히 실행된다
// 가장 먼저 접근한 고루틴만 락을 걸어 혼자만 사용할 수 있고 나머지는 대기한다
package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	account := &Account{10}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			// 무한루프
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// 깃발의 되는 곳
func DepositAndWithdraw(account *Account) {
	// 락(Lock)
	mutex.Lock()
	// 언락(Unlock)
	defer mutex.Unlock()

	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be nagative value: %d", account.Balance))
	}
	// 10개의 고루틴이 실행되는 동안 1000원을 플러스하고 마이너스 하는 동작을 무한히 반복
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}
