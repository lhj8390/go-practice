package main

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

type Message struct {
	Name      string
	Content   string
	Timestamp string
}

func main() {
	// 기본 nsq 구성 설정
	config := nsq.NewConfig()
	// NSQD 주소를 사용하여 Producer 인스턴스 반환
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}
	// 초기 topic 이름 및 메시지 구성
	topic := "Topic_Example"
	msg := Message{
		Name:      "Message Name Example",
		Content:   "Message Content Example",
		Timestamp: time.Now().String(),
	}
	// json 객체 마샬링
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	// message publish
	err = producer.Publish(topic, payload)
	if err != nil {
		log.Println(err)
	}
}
