package config

import (
	"fmt"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"serviceOpname-v2/config/entity"
)

func SetUpDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	//regis db
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	//connect db
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a conn")
	}

	// make model in here
	db.AutoMigrate(&entity.User{}, &entity.Opname{})
	return db
}

// close db
func CloseConnDb(db *gorm.DB){
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close conn db")
	}
	dbSQL.Close()
}