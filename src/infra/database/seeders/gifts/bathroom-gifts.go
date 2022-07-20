package gifts_seeder

import (
	"go_gift_list_api/src/infra/database/models"
	"log"

	"gorm.io/gorm"
)

func (g *GiftSeeder) createBathroomGifts(db *gorm.DB) {
	var bathroomCategoryID int64 = 3
	var bathroomGifts = []models.Gift{
		{
			Name:       "Chuveiro Elétrico",
			CategoryID: bathroomCategoryID,
			Available:  1,
			Image:      "chuveiro",
			Quantity:   1,
			PriceGrade: 3,
		},
		{
			Name:       "Jogo de tapetes",
			CategoryID: bathroomCategoryID,
			Available:  1,
			Image:      "jogo_tapetes_banheiro",
			Quantity:   1,
			PriceGrade: 1,
		},
		{
			Name:       "Kit de acessórios",
			CategoryID: bathroomCategoryID,
			Available:  1,
			Image:      "kit_acessorios_banheiro",
			Quantity:   1,
			PriceGrade: 2,
		},
		{
			Name:       "Lixeira de banheiro",
			CategoryID: bathroomCategoryID,
			Available:  1,
			Image:      "lixeira_de_banheiro",
			Quantity:   1,
			PriceGrade: 2,
		},
		{
			Name:       "Toalhas de banho",
			CategoryID: bathroomCategoryID,
			Available:  1,
			Image:      "toalhas_de_banho",
			Quantity:   1,
			PriceGrade: 1,
		},
		{
			Name:       "Toalhas de rosto",
			CategoryID: bathroomCategoryID,
			Available:  1,
			Image:      "toalhas_de_rosto",
			Quantity:   1,
			PriceGrade: 1,
		},
	}

	for i := range bathroomGifts {
		err := db.Debug().Model(&models.Gift{}).Create(&bathroomGifts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed bathroom gifts table: %v", err)
		}
	}
}
