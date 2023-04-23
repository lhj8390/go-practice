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

	account := &Account{0}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock() // 뮤텍스를 사용하여 동시성 문제 해결
	// 한 번 획득한 뮤텍스는 반드시 Unlock()을 호출해서 반납해야 한다.
	// 뮤텍스 단점 : 성능 향상을 얻을 수 없다, 데드락이 발생할 수 있다.
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d",
			account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}
