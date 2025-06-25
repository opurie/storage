package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Description string `json:"description"`
}

type Item struct {
	ID int
	Name string
	Category Category
	Location Location
	Owner User
	Description string
	Tags []string
}

type Image struct {
	ID int
	ItemID int
	Path string
}

type Category struct {
	ID int
	Name string
}
type Location struct {
	ID int
	Name string
}


func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Item{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Location{})
	db.AutoMigrate(&Image{})

	// Add foreign keys
	db.Model(&Item{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	db.Model(&Item{}).AddForeignKey("location_id", "locations(id)", "RESTRICT", "RESTRICT")
	db.Model(&Item{}).AddForeignKey("owner_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&Image{}).AddForeignKey("item_id", "items(id)", "RESTRICT", "RESTRICT")	
	// Add secondary keys
	db.Model(&User{}).AddUniqueIndex("idx_username", "username")
	return db
}