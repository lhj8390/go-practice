package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Book struct {
	gorm.Model
	ID   int `gorm:"primaryKey"`
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect DB : %v", err)
	}
	// migration
	db.AutoMigrate(&Book{})

	db.Create(&Book{ID: 1, Name: "test"})

	// 데이터 조회
	var book Book
	db.First(&book)
	fmt.Printf("result : (%v %v)", book.ID, book.Name)
}
