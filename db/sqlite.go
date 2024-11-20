package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
}

func (Product) TableName() string {
	return "products"
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db/sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic("failed to migrate database schema")
	}

	return db
}
