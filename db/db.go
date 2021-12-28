package db

import (
	"github.com/jinzhu/gorm"
	"master_go_programming/models"
	"os"
)

var db *gorm.DB
var err error

// Init gorm db
func Init() *gorm.DB {
	DRIVE := "mysql"
	HOST := os.Getenv("DATABASE_HOST")
	PORT := os.Getenv("DATABASE_PORT")
	USER := os.Getenv("DATABASE_USER")
	PASS := os.Getenv("DATABASE_PASSWORD")
	DBNAME := os.Getenv("DATABASE_NAME")

	CONNECT := USER + ":" + PASS + "@" + "tcp(" + HOST + ":" + PORT + ")" + "/" + DBNAME + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true"

	db, err = gorm.Open(DRIVE, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	if os.Getenv("ENVIRONMENT") == "development" {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Post{})
	} else {
		db.Set("gorm:table_options", "ENGINE=InnoDB")
		db.LogMode(true)
	}

	return db
}

// DbManager get db connector
func DbManager() *gorm.DB {
	return db
}
