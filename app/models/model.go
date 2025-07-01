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
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
	LocationID  uint
	Location    Location `gorm:"foreignKey:LocationID"`
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	Description string
	Tags        string
}

type Image struct {
	gorm.Model

	ItemID uint
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
	db.LogMode(true)
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Location{})
	db.AutoMigrate(&Item{})
	// db.AutoMigrate(&Image{})

	db.Model(&Item{}).AddForeignKey(
		"category_id", // kolumna w Item
		"categories(id)", // tabela i kolumna docelowa
		"RESTRICT",  // ON DELETE
		"RESTRICT",  // ON UPDATE
	)
	db.Model(&Item{}).AddForeignKey("location_id", "locations(id)", "RESTRICT", "RESTRICT")
	db.Model(&Item{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	return db
}
