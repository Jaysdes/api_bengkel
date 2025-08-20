package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	user := os.Getenv("MYSQLUSER")
	pass := os.Getenv("MYSQLPASSWORD")
	host := os.Getenv("MYSQLHOST")
	port := os.Getenv("MYSQLPORT")
	dbname := os.Getenv("MYSQLDATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s",
		user, pass, host, port, dbname)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect database:", err)
	}

	DB = database
	fmt.Println("✅ Database connected to", dbname)
}
