package migrations

import (
	"errors"
	"go_gift_list_api/src/infra/database/models"
	categories_seeder "go_gift_list_api/src/infra/database/seeders/categories"
	gifts_seeder "go_gift_list_api/src/infra/database/seeders/gifts"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Category{}); err == nil && db.Migrator().HasTable(&models.Category{}) {
		if err := db.First(&models.Category{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			categorySeeder := &categories_seeder.CategorySeeder{}
			categorySeeder.Run(db)
		}
	}

	if err := db.AutoMigrate(&models.Gift{}); err == nil && db.Migrator().HasTable(&models.Gift{}) {
		if err := db.First(&models.Gift{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			giftSeeder := &gifts_seeder.GiftSeeder{}
			giftSeeder.Run(db)
		}
	}
}
