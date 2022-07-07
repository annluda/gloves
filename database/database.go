package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gloves/pkg/logger"
)

// DB gorm
var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "gorm.db")
	if err != nil {
		logger.Fatalf("Database connection failed. error: ", err)
	}

	//db.LogMode(config.DBConfig.Debug)
	DB = db

	return db
}
