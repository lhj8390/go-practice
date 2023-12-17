package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model // ID, CreatedAt, UpdatedAt, DeletedAt 필드를 포함하는 구조체
	Name       string
	AuthorId   int
}

type Author struct {
	ID   int
	Name string
}

// ================================
// Eager Loading Example
// ================================

type User struct {
	Id        uint
	Name      string
	Orders    []Order
	CompanyId uint
	Company   Company
}

type Order struct {
	Id     uint
	Price  int
	UserId uint
}

type Company struct {
	Id   uint
	Name string
}
