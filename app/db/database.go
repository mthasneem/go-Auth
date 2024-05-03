package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() {
	// MySQL connection details
	dsn := "root@tcp(127.0.0.1:3306)/go_basic?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to the MySQL database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db

	// AutoMigrate the models
	db.AutoMigrate(&User{})
}
