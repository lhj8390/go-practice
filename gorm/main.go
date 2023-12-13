package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Book struct {
	gorm.Model // ID, CreatedAt, UpdatedAt, DeletedAt 필드를 포함하는 구조체
	Name       string
}

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect DB : %v", err)
	}

	defer func() {
		// 전체 코드가 완전히 실행되면 연결을 종료한다.
		// GORM v1.20 부터는 connection pooling 을 지원하기 때문에
		// connection 을 열고 어플리케이션 내에서 공유하는 것이 적절한 방법.
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()
	// migration
	err = db.AutoMigrate(&Book{})
	if err != nil {
		return
	}

	// 데이터 조회
	var books []Book
	db.Find(&books)
	fmt.Printf("result : (%v)", len(books))
}
