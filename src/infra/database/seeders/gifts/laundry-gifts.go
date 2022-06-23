package gifts_seeder

import (
	"go_gift_list_api/src/infra/database/models"
	"log"

	"gorm.io/gorm"
)

func (g *GiftSeeder) createLaundryGifts(db *gorm.DB) {
	var laundryCategoryID int64 = 3
	var laundryGifts = []models.Gift{
		{
			Name:          "Lava e seca",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "lava_e_seca",
			Quantity:      1,
			PriceGrade:    3,
		},
		{
			Name:          "Aspirador de pó",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "aspirador_de_po",
			Quantity:      1,
			PriceGrade:    3,
		},
		{
			Name:          "Varal de chao",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "varal_de_chao",
			Quantity:      1,
			PriceGrade:    2,
		},
		{
			Name:          "Varal oculto",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "varal_oculto",
			Quantity:      1,
			PriceGrade:    2,
		},
		{
			Name:          "Vassoura",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "vassoura",
			Quantity:      2,
			PriceGrade:    1,
		},
		{
			Name:          "Rodo",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "rodo",
			Quantity:      2,
			PriceGrade:    1,
		},
		{
			Name:          "Balde",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "balde",
			Quantity:      2,
			PriceGrade:    1,
		},
		{
			Name:          "Pá",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "pa",
			Quantity:      1,
			PriceGrade:    1,
		},
		{
			Name:          "Ferro de passar roupa",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "ferro_de_passar_roupa",
			Quantity:      1,
			PriceGrade:    2,
		},
		{
			Name:          "Tábua de passar roupa",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "tabua_de_passar_roupa",
			Quantity:      1,
			PriceGrade:    2,
		},
		{
			Name:          "Cesto de roupas",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "cesto_de_roupas",
			Quantity:      1,
			PriceGrade:    2,
		},
		{
			Name:          "Prendedores de roupa",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "prendedores_de_roupa",
			Quantity:      1,
			PriceGrade:    1,
		},
		{
			Name:          "Cabides de roupa",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "prendedores_de_roupa",
			Quantity:      2,
			PriceGrade:    1,
		},
		{
			Name:          "Panos de chão",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "panos_de_chao",
			Quantity:      2,
			PriceGrade:    1,
		},
		{
			Name:          "Panos de chão",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "panos_de_chao",
			Quantity:      2,
			PriceGrade:    1,
		},
		{
			Name:          "Flanelas",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "flanelas",
			Quantity:      2,
			PriceGrade:    1,
		},
		{
			Name:          "Espanador de pó",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "espanador_de_po",
			Quantity:      2,
			PriceGrade:    1,
		},
		{
			Name:          "Caixa organizadora",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "caixa_organizadora",
			Quantity:      2,
			PriceGrade:    1,
		},
		{
			Name:          "Escada",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "escada",
			Quantity:      1,
			PriceGrade:    2,
		},
		{
			Name:          "Kit de costura",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "kit_de_costura",
			Quantity:      1,
			PriceGrade:    1,
		},
		{
			Name:          "Kit básico de ferramentas",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "kit_basico_de_ferramentas",
			Quantity:      1,
			PriceGrade:    2,
		},
		{
			Name:          "Caixa primeiros socorros",
			CategoryID:    laundryCategoryID,
			Available:     1,
			ImageFilename: "caixa_primeiros_socorros",
			Quantity:      1,
			PriceGrade:    1,
		},
	}

	for i := range laundryGifts {
		err := db.Debug().Model(&models.Gift{}).Create(&laundryGifts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed laundry gifts table: %v", err)
		}
	}
}
