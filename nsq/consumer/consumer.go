package main

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type messageHandler struct{}
type Message struct {
	Name      string
	Content   string
	Timestamp string
}

func main() {
	// nsq 의 config 인스턴스 생성 및 속성 정의
	config := nsq.NewConfig()

	// 최대 메시지 처리 시도 횟수
	config.MaxAttempts = 10
	// 최대 동시 처리 메시지 수
	config.MaxInFlight = 5
	// 메시지가 재 처리될 때 최대, 기본 대기 시간
	config.MaxRequeueDelay = time.Second * 900
	config.DefaultRequeueDelay = time.Second * 0

	// topic 이름 및 채널 초기화
	topic := "Topic_Example"
	channel := "Channel_Example"
	// consumer 생성
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}

	// handler 등록
	// 메시지를 처리하는 로직이 들어가 있다
	consumer.AddHandler(&messageHandler{})
	// NSQLookupd를 통해 nsqd 인스턴스 찾기
	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Fatal(err)
	}
	// SIGINT 또는 SIGTERM 시그널 기다리기
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	// consumer Gracefully 하게 종료
	consumer.Stop()
}

// HandleMessage nsq 패키지의 Handler interface 상속
func (h *messageHandler) HandleMessage(m *nsq.Message) error {

	var request Message
	if err := json.Unmarshal(m.Body, &request); err != nil {
		log.Println("Error when Unmarshaling the message body, Err : ", err)
		// 오류가 아닌 값을 반환하면 성공으로 간주되지 않아 NSQ는 해당 메시지를 다시 큐에 넣어 재처리한다
		return err
	}
	//Print the Message
	log.Println("Message")
	log.Println("--------------------")
	log.Println("Name : ", request.Name)
	log.Println("Content : ", request.Content)
	log.Println("Timestamp : ", request.Timestamp)
	log.Println("--------------------")
	log.Println("")

	// 메시지를 성공적으로 처리했으므로 nil 을 반환
	// NSQ 는 성공되었다고 인식하여 해당 메시지를 큐에서 삭제
	return nil
}
