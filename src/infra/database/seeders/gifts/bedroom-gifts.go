package gifts_seeder

import (
	"go_gift_list_api/src/infra/database/models"
	"log"

	"gorm.io/gorm"
)

func (g *GiftSeeder) createBedroomGifts(db *gorm.DB) {
	var bedroomCategoryID int64 = 4
	var bedroomGifts = []models.Gift{
		{
			Name:       "Espelho de corpo inteiro",
			CategoryID: bedroomCategoryID,
			Available:  1,
			Image:      "espelho_corpo_inteiro",
			Quantity:   1,
			PriceGrade: 3,
		},
		{
			Name:       "Smart TV 32 polegadas",
			CategoryID: bedroomCategoryID,
			Available:  1,
			Image:      "tv_32_polegadas",
			Quantity:   1,
			PriceGrade: 3,
		},
		{
			Name:       "Cortinas",
			CategoryID: bedroomCategoryID,
			Available:  1,
			Image:      "cortinas",
			Quantity:   1,
			PriceGrade: 2,
		},
		{
			Name:       "Cabides de roupa",
			CategoryID: bedroomCategoryID,
			Available:  1,
			Image:      "cabides_de_roupa",
			Quantity:   1,
			PriceGrade: 2,
		},
	}

	for i := range bedroomGifts {
		err := db.Debug().Model(&models.Gift{}).Create(&bedroomGifts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed bedroom gifts table: %v", err)
		}
	}
}
