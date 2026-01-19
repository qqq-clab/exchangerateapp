package config

import (
	"exchangeapp/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() {
	dsn := AppConfig.Database.Dsn
	// log.Fatalf("%v", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db *gorm.DB

	if err != nil {
		log.Fatalf("Failed to initial database with error code : %v", err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatalf("Failed to config database with error code : %v", err)
	}
	global.Db = db
}
