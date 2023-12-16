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
	Id     int
	Name   string
	Orders []Order
}

type Order struct {
	Id    int
	Price int
}
