package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lexkong/log"
)

// DB gorm
var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		log.Fatal("Database connection failed. error: ", err)
	} else {
		fmt.Print("\n------------ GORM OPEN SUCCESS! --------------\n")
	}

	//db.LogMode(config.DBConfig.Debug)
	DB = db

	return db
}
