package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"unique_index"`
	Email       string `json:"email"`
	Description string `json:"description"`
}


type Item struct {
	gorm.Model
	Name        string
	Category    Category `gorm:"embedded`
	Location    Location `gorm:"embedded"`
	Owner       User     `gorm:"embedded"`
	Description string
	Tags        string
}

type Image struct {
	gorm.Model
	ItemID string
	Path   string
}

type Category struct {
	gorm.Model
	Name string
}
type Location struct {
	gorm.Model
	Name string
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Item{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Location{})
	db.AutoMigrate(&Image{})

	return db
}
