package query

import (
	"fmt"
	"go-practice/gorm/model"
	"gorm.io/gorm"
)

func SelectOne(db *gorm.DB) {
	var book model.Book
	db.First(&book, 1)
	fmt.Println("========== SELECT ONE ===========")
	fmt.Printf("result id=1 : (%v) \n", book.ID)

	book2 := model.Book{Name: "book2"}
	db.First(&book)
	fmt.Printf("result name=book2 : (%v)", book2.Name)
}

func SelectAll(db *gorm.DB) {
	var books []model.Book
	db.Find(&books)
	fmt.Println("========== SELECT ALL ===========")
	fmt.Printf("result : (%v)", len(books))
}
