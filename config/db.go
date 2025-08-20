package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// DSN format: username:password@tcp(host:port)/dbname?parseTime=True
	dsn := "lesc6121:Ykx578uUzb2t25ingin@tcp(bangli.iixcp.rumahweb.net:3306)/lesc6121_bengkel?parseTime=True"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	DB = database
	fmt.Println("Database connected successfully")
}
