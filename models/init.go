package models

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func sqliteDB(dsn string, config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn+"?_pragma=foreign_keys(1)"), config)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(1)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB() {
	var db *gorm.DB
	var err error
	db, err = sqliteDB("weixin_DB.db", &gorm.Config{})
	if err != nil {
		log.Panicf("无法连接数据库，%s", err)
	}

	DB = db
	Migrate()
}

func Migrate() {
	// DB.AutoMigrate(&User{})
}
