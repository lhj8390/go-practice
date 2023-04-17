package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

// 포인트 메서드
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 값 타입 메서드
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

// 변경된 값을 반환하는 값 타입 메서드
func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	var mainA = &account{100, "Joe", "Park"}
	mainA.withdrawPointer(30)
	fmt.Println("포인트 메서드 (-30) : ", mainA.balance)

	mainA.withdrawValue(20) // (*mainA) 값 타입으로 자동 변환
	fmt.Println("값 타입 메서드 (-20) : ", mainA.balance)

	var mainB = mainA.withdrawReturnValue(20)
	fmt.Println("변경된 값을 반환하는 값 타입 메서드 (-20) : ", mainB.balance)

	mainB.withdrawPointer(30) // (&mainB) 포인터로 자동 변환
	fmt.Println("포인트 메서드 (-30) : ", mainB.balance)
}
