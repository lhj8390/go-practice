package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var num int
	fmt.Print("숫자값을 입력하세요:")
	_, err := fmt.Scanln(&num)
	if err != nil {
		stdin.ReadString('\n')
	}
	return num, err
}

func GenerateRan(n int) int {
	source := rand.NewSource(time.Now().UnixNano())
	ran := rand.New(source)
	return ran.Intn(n)
}

func main() {
	count := 0
	n := GenerateRan(100)
	for {
		num, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
			continue
		}

		if num > n {
			fmt.Println("입력하신 숫자가 더 큽니다.")
		} else if num < n {
			fmt.Println("입력하신 숫자가 더 작습니다.")
		} else {
			fmt.Print("숫자를 맞췄습니다. 축하합니다. 시도한 횟수: ")
			fmt.Println(count)
			break
		}
		count++
	}
}
