package gifts_seeder

import (
	"go_gift_list_api/src/infra/database/models"
	"log"

	"gorm.io/gorm"
)

func (g *GiftSeeder) createNewGifts(db *gorm.DB) {
	var laundryCategoryID int64 = 3
	var bedroomCategoryID int64 = 4
	var livingRoomCategoryID int64 = 5
	var outherCategoryID int64 = 6

	var gifts = []models.Gift{
		{
			Name:       "Kit limpeza",
			CategoryID: laundryCategoryID,
			Available:  1,
			Image:      "kit_limpeza",
			Quantity:   2,
			PriceGrade: 1,
		},
		{
			Name:       "Porta retrato",
			CategoryID: livingRoomCategoryID,
			Available:  1,
			Image:      "porta_retrato",
			Quantity:   5,
			PriceGrade: 1,
		},
		{
			Name:       "Jogo de cama",
			CategoryID: bedroomCategoryID,
			Available:  1,
			Image:      "jogo_de_cama",
			Quantity:   4,
			PriceGrade: 1,
		},
		{
			Name:       "Planta decorativa",
			CategoryID: outherCategoryID,
			Available:  1,
			Image:      "planta decorativa",
			Quantity:   4,
			PriceGrade: 1,
		},
	}

	for i := range gifts {
		err := db.Debug().Model(&models.Gift{}).Create(&gifts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed laundry gifts table: %v", err)
		}
	}
}
