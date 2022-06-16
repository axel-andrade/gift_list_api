package gifts_seeder

import (
	"go_gift_list_api/src/infra/database/models"
	"log"

	"gorm.io/gorm"
)

func (g *GiftSeeder) createBathroomGifts(db *gorm.DB) {
	var bathroomCategoryID int64 = 1
	var bathroomGifts = []models.Gift{
		{
			Name:          "Chuveiro Elétrico",
			CategoryID:    bathroomCategoryID,
			Available:     1,
			ImageFilename: "chuveiro",
			Quantity:      1,
		},
		{
			Name:          "Jogo de tapetes",
			CategoryID:    bathroomCategoryID,
			Available:     1,
			ImageFilename: "jogo_tapetes_banheiro",
			Quantity:      1,
		},
		{
			Name:          "Kit de acessórios",
			CategoryID:    bathroomCategoryID,
			Available:     1,
			ImageFilename: "kit_acessorios_banheiro",
			Quantity:      1,
		},
		{
			Name:          "Kit de acessórios",
			CategoryID:    bathroomCategoryID,
			Available:     1,
			ImageFilename: "kit_acessorios_banheiro",
			Quantity:      1,
		},
		{
			Name:          "Toalhas de banho",
			CategoryID:    bathroomCategoryID,
			Available:     1,
			ImageFilename: "toalhas_de_banho",
			Quantity:      1,
		},
		{
			Name:          "Toalhas de rosto",
			CategoryID:    bathroomCategoryID,
			Available:     1,
			ImageFilename: "toalhas_de_rosto",
			Quantity:      1,
		},
	}

	for i := range bathroomGifts {
		err := db.Debug().Model(&models.Gift{}).Create(&bathroomGifts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed bathroom gifts table: %v", err)
		}
	}
}
