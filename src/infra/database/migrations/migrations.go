package migrations

import (
	"go_gift_list_api/src/infra/database/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Category{})
}
