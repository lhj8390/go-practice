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

// Preload Eager Loading (즉시 로딩) 예제
func Preload(db *gorm.DB) {
	var users []model.User

	// SELECT * FROM `orders` WHERE `orders`.`user_id` IN (?,?,?, ....?,?)
	// SELECT * FROM `users`
	db.Preload("Orders").Find(&users)

	fmt.Println("========== PRELOAD ===========")
	fmt.Printf("User: %s\n", users[0].Name)
	fmt.Printf("Orders : %d", len(users[0].Orders))
}

// LazyLoading 기본 조회 예제
func LazyLoading(db *gorm.DB) {
	var users []model.User

	// SELECT * FROM `users`
	db.Find(&users)

	fmt.Println("========== LAZY LOADING ===========")
	fmt.Printf("User: %s\n", users[0].Name)
	fmt.Printf("Orders is nil : %v\n", len(users[0].Orders)) // Lazy Loading 이기 때문에 연관 테이블 order 데이터 누락.
}

func Join(db *gorm.DB) {
	var users []model.User

	// SELECT `users`.`id`,`users`.`name`,`users`.`company_id`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users`
	// LEFT JOIN `companies` `Company` ON `users`.`company_id` = `Company`.`id`
	db.Joins("Company").Find(&users)

	fmt.Println("========== JOIN (PRELOADING) ===========")
	fmt.Println("Joins Preloading 은 1:1 관계일 경우에만 사용할 수 있다.")
	fmt.Printf("User: %s\n", users[0].Name)
	fmt.Printf("User[0]'s Company : (%v, %v)\n", users[0].Company.Id, users[0].Company.Name)

}
