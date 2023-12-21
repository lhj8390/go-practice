package main

import (
	"fmt"
	"go-practice/gorm/model"
	"go-practice/gorm/query"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: getLogger(),
	})
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

	// 테스트를 위해 테이블을 drop 한다.
	err = db.Migrator().DropTable(&model.Book{}, &model.Author{}, &model.Order{}, &model.User{}, &model.Company{})
	if err != nil {
		log.Fatalf("Failed to drop table: %v", err)
	}

	// migration
	err = db.AutoMigrate(&model.Book{}, &model.Author{}, &model.Order{}, &model.User{}, &model.Company{})
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}

	// 테스트 데이터 생성
	generateData(db)

	// 데이터 단건 조회
	query.SelectOne(db)
	// 데이터 전체 조회
	query.SelectAll(db)

	// preload
	query.Preload(db)
	// lazy loading (default)
	query.LazyLoading(db)
	// Join (Preloading)
	query.Join(db)

	// Associations
	// find associations
	query.FindAssociations(db)
	// append associations
	query.AppendAssociations(db)
	// replace associations
	query.ReplaceAssociations(db)
}

func generateData(db *gorm.DB) {
	var author []model.Author
	var book []model.Book
	var users []model.User

	for i := range [100]int{} {
		author = append(author, model.Author{
			Name: fmt.Sprintf("author%v", i),
		})
		book = append(book, model.Book{
			AuthorId: i,
			Name:     fmt.Sprintf("book%v", i),
		})
		users = append(users, model.User{
			Name: fmt.Sprintf("user%v", i),
			Orders: []model.Order{
				{
					Price: (i + 1) * 1000,
				},
				{
					Price: (i + 2) * 1000,
				},
			},
			Company: model.Company{
				Id:   uint(i),
				Name: fmt.Sprintf("company%v", i),
			},
		})
	}

	// batch insert
	db.CreateInBatches(author, 100)
	db.CreateInBatches(book, 100)
	db.Create(&users)
}

func getLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:        time.Second, // Slow SQL threshold
			LogLevel:             logger.Info, // Log level
			ParameterizedQueries: true,        // Don't include params in the SQL log
			Colorful:             true,        // Disable color
		},
	)
	return newLogger
}
