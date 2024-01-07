package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	db_name := os.Getenv("db_name")
	db_host := os.Getenv("db_host")
	db_port := os.Getenv("db_port")
	db_username := os.Getenv("db_username")
	db_password := os.Getenv("db_password")

	log.Println(db_name)
	var err error
	dsn := db_username + ":" + db_password + "@tcp" + "(" + db_host + ":" + db_port + ")/" + db_name + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
