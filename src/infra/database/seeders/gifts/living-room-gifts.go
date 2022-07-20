package gifts_seeder

import (
	"go_gift_list_api/src/infra/database/models"
	"log"

	"gorm.io/gorm"
)

func (g *GiftSeeder) createlivingRoomGifts(db *gorm.DB) {
	var livingRoomCategoryID int64 = 5
	var livingRoomGifts = []models.Gift{
		{
			Name:       "Sofá cama",
			CategoryID: livingRoomCategoryID,
			Available:  1,
			Image:      "sofa_cama",
			Quantity:   1,
			PriceGrade: 3,
		},
		{
			Name:       "Almofadas",
			CategoryID: livingRoomCategoryID,
			Available:  1,
			Image:      "almofadas",
			Quantity:   1,
			PriceGrade: 2,
		},
		{
			Name:       "Ventilador de coluna",
			CategoryID: livingRoomCategoryID,
			Available:  1,
			Image:      "ventilador_de_coluna",
			Quantity:   1,
			PriceGrade: 3,
		},
		{
			Name:       "Climatizador",
			CategoryID: livingRoomCategoryID,
			Available:  1,
			Image:      "climatizador",
			Quantity:   1,
			PriceGrade: 3,
		},
		{
			Name:       "Porta chaves",
			CategoryID: livingRoomCategoryID,
			Available:  1,
			Image:      "porta_chaves",
			Quantity:   1,
			PriceGrade: 3,
		},
		{
			Name:       "Tapete porta de entrada",
			CategoryID: livingRoomCategoryID,
			Available:  1,
			Image:      "tapete_porta_de_entrada",
			Quantity:   1,
			PriceGrade: 1,
		},
	}

	for i := range livingRoomGifts {
		err := db.Debug().Model(&models.Gift{}).Create(&livingRoomGifts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed living room gifts table: %v", err)
		}
	}
}
