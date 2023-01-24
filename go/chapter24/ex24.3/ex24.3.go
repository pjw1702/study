package main

import (
	"fmt"
	"sync"
	"time"
)

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

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be nagative value: %d", account.Balance))
	}
	// 10개의 고루틴이 실행되는 동안 1000원을 플러스하고 마이너스 하는 동작을 무한히 반복
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

/*
* 1000을 계속 더하고 빼는 프로그램인데도 'account.Balance'의 필드 값이 0보다 적은 수가 입력되어 패닉이 발생하여 프로그램이 종료된다
	* 고루틴의 동시성 문제
	* 우리 눈에는 account.Balance += 1000 명령어 한 줄이지만, CPU는 ADD account.Balance 1000, MOV account.Balance Rx1의 두 개의 명령어가 레지스터에 적재되어 있다
		* 여러 개의 고루틴이 동시에 ADD 명령과 MOV 명령에 걸리게 되면, 어떤 고루틴 들은 ADD 명령에 걸리고 어떤 고루틴들은 MOV 명령에 고루틴에 걸리게 되면서 값이 꼬이게 된다
	* 하나의 메모리 자원의 여러 고루틴이 접근하기 때문에 생긴 문제
		* 막아줘야 한다
*/
