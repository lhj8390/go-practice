package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"runtime"
	"time"
)

func main() {

	var (
		maxWorker = runtime.GOMAXPROCS(0)
		sem       = semaphore.NewWeighted(int64(maxWorker))
		total     = 100
	)

	// 동시에 실행할 수 있는 워커 개수 제한
	fmt.Printf("Running with %d workers\n", maxWorker)
	ctx := context.Background()
	for i := 0; i <= total; i++ {
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}

		go func(v int) {
			defer sem.Release(1)
			runningProcess(v)
		}(i)
	}

	// 슬롯이 모두 해제될 때까지 대기
	if err := sem.Acquire(ctx, int64(maxWorker)); err != nil {
		log.Printf("Failed to acquire semaphore: %v\n", err)
	}
}

func runningProcess(taskId int) {
	time.Sleep(2 * time.Second)
	log.Printf(" Running task with ID : %v", taskId)

}
