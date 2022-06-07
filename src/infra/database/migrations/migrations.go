package migrations

import (
	"errors"
	"go_gift_list_api/src/infra/database/models"
	"go_gift_list_api/src/infra/database/seeders"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {

	if err := db.AutoMigrate(&models.Category{}); err == nil && db.Migrator().HasTable(&models.Category{}) {
		if err := db.First(&models.Category{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			categorySeeder := &seeders.CategorySeeder{}
			categorySeeder.Run(db)
		}
	}
}
