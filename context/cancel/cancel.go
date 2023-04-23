package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 작업이 취소될 때마다 1초 간격으로 메시지 출력
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background()) // 취소 가능한 컨텍스트 생성
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //3초 시간 동안만 작업 지시 가능
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel() // 취소함수 (원할 때 취소 가능)

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}

	}
}
