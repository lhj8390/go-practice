package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	// 각 고루틴이 같은 자원에 접근하지 않도록 하는 방법
	// 1. 영역을 나누는 방법
	// 2. 역할을 나누는 방법

	// 2. 역할을 나누는 방법 예제 (소비자 패턴)
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}

// MakeBody 1초 간격으로 car 생성 후 tireCh 에 데이터 전달
func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			car := &Car{}
			car.Body = "Sports car"
			tireCh <- car
		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

// InstallTire tireCh 에서 데이터를 읽고 paintCh 채널에 전달
func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "Winter tire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

// PaintCar paintCh 채널에서 데이터를 읽고 출력. 채널이 닫히면 루틴 종료
func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(),
			car.Body, car.Tire, car.Color)
	}
	wg.Done()
}
