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

// FindAssociations 연관된 결과 찾기
func FindAssociations(db *gorm.DB) {
	var orders []model.Order

	user := model.User{Id: 1, Name: "user0"} // id 가 1 인 User 와 연관된 Order 찾기
	// SELECT * FROM `orders` WHERE `orders`.`user_id` = ?
	db.Model(&user).Association("Orders").Find(&orders)

	fmt.Println("========== FIND ASSOCIATIONS ===========")
	fmt.Printf("User: (%v, %v)\n", user.Id, user.Name)
	fmt.Printf("Orders[0]: (%v, %v)\n", orders[0].Id, orders[0].Price)
}

// AppendAssociations (N:N, 1:N 관계일 경우) 새로운 연관관계 추가, (1:1 관계일 경우) 연관관계 교체
func AppendAssociations(db *gorm.DB) {

	user := model.User{Id: 1, Name: "user0"}
	// INSERT INTO `orders` (`price`,`user_id`,`id`) VALUES (?,?,?),,(?,?,?),(?,?,?)
	// ON CONFLICT (`id`) DO UPDATE SET `user_id`=`excluded`.`user_id` RETURNING `id`
	db.Model(&user).Association("Orders").Append([]model.Order{
		{Id: 2, Price: 2000, UserId: 1},
		{Id: 4, Price: 4000, UserId: 1},
		{Id: 6, Price: 6000, UserId: 1},
	})

	// 기존 DB의 값 : Order(1, 1000), Order(2, 2000)
	// 추가된 값 : Order(2, 2000), Order(4, 4000), Order(6, 6000)  -> 테이블의 연관관계가 변경된다.
	fmt.Println("========== APPEND ASSOCIATIONS ===========")
	fmt.Printf("User: (%v, %v)\n", user.Id, user.Name)
	fmt.Printf("Orders[0], Orders[1], Orders[2]: (%v, %v, %v)\n", user.Orders[0].Price, user.Orders[1].Price, user.Orders[2].Price)
}

// ReplaceAssociations 새로운 연관관계로 교체 (1:1 관계의 Append Associations 와 동일)
func ReplaceAssociations(db *gorm.DB) {
	user := model.User{Id: 1, Name: "user0"}
	// INSERT INTO `companies` (`name`,`id`) VALUES (?,?) ON CONFLICT DO NOTHING RETURNING `id`
	// UPDATE `users` SET `company_id`=? WHERE `id` = ?
	db.Model(&user).Association("Company").Replace(&model.Company{Id: 3, Name: "company22"})

	// 기존 DB의 값 : Company(1, company0)
	// 추가된 값 : Company(3, company22)
	// company_id 가 3인 행이 이미 존재하므로 기본이 되는 테이블 (User)의 company_id 만 변경된다.
	fmt.Println("========== REPLACE ASSOCIATIONS ===========")
	fmt.Printf("User: (%v, %v)\n", user.Id, user.Name)
	fmt.Printf("Company: (%v)\n", user.Company)
}
