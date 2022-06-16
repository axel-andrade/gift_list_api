package categories_seeder

import (
	"go_gift_list_api/src/infra/database/models"
	"log"

	"gorm.io/gorm"
)

type CategorySeeder struct{}

func (c *CategorySeeder) Run(db *gorm.DB) {
	var categories = []models.Category{
		{
			Name: "Banheiro",
		},
		{
			Name: "Cozinha",
		},
		{
			Name: "Lavanderia",
		},
		{
			Name: "Quarto",
		},
		{
			Name: "Sala de estar",
		},
		{
			Name: "Outros",
		},
	}

	for i, _ := range categories {
		err := db.Debug().Model(&models.Category{}).Create(&categories[i]).Error
		if err != nil {
			log.Fatalf("cannot seed categories table: %v", err)
		}
	}
}
