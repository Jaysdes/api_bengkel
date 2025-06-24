package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(localhost)/bengkel?parseTime=True"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}

	DB = database

	fmt.Println("Database connected ")
}
