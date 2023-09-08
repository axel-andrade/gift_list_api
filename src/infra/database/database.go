package database

import (
	"go_gift_list_api/src/infra/database/migrations"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(1)
	config.SetMaxOpenConns(1)
	config.SetConnMaxLifetime(time.Hour)

	if os.Getenv("DB_AUTO_MIGRATE") == "true" {
		migrations.Run(db)
	}

	// Com o defer o go vai conseguir identificar quando executar uma determinada ação
	// defer config.Close()
}

func CloseDB() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
