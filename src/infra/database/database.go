package database

import (
	"go_gift_list_api/src/infra/database/migrations"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	if os.Getenv("DB_AUTO_MIGRATE") == "true" {
		migrations.RunMigrations(db)
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
