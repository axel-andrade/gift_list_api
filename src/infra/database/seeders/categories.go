package seeders

import (
	"go_gift_list_api/src/infra/database/models"
	"log"
	"time"

	"gorm.io/gorm"
)

type CategorySeeder struct{}

func (c *CategorySeeder) Run(db *gorm.DB) {
	var categories = []models.Category{
		{
			Base: models.Base{
				ID:        1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Banheiro",
		},
		{
			Base: models.Base{
				ID:        2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Cozinha",
		},
		{
			Base: models.Base{
				ID:        3,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Lavanderia",
		},
		{
			Base: models.Base{
				ID:        4,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Quarto",
		},
		{
			Base: models.Base{
				ID:        5,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: "Sala de estar",
		},
		{
			Base: models.Base{
				ID:        6,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
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
