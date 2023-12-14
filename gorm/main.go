package main

import (
	"fmt"
	"go-practice/gorm/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

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

	// migration 전에 수행할 작업
	// 테스트를 위해 테이블을 drop 한다.
	err = db.Migrator().DropTable(&model.Book{}, &model.Author{})
	if err != nil {
		log.Fatalf("Failed to drop table: %v", err)
	}

	// migration
	err = db.AutoMigrate(&model.Book{}, &model.Author{})
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	// 테스트 데이터 생성
	generateData(db)

	// 데이터 조회
	var books []model.Book
	db.Find(&books)
	fmt.Printf("result : (%v)", len(books))
}

func generateData(db *gorm.DB) {
	var author []model.Author
	var book []model.Book

	for i := range [100]int{} {
		author = append(author, model.Author{
			Name: fmt.Sprintf("author%v", i),
		})
		book = append(book, model.Book{
			AuthorId: i,
			Name:     fmt.Sprintf("book%v", i),
		})
	}

	// batch insert
	db.CreateInBatches(author, 100)
	db.CreateInBatches(book, 100)
}
