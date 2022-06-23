package gifts_seeder

import (
	"gorm.io/gorm"
)

type GiftSeeder struct{}

func (g *GiftSeeder) Run(db *gorm.DB) {
	g.createBathroomGifts(db)
	g.createKitchenGifts(db)
	g.createLaundryGifts(db)
	g.createBedroomGifts(db)
	g.createlivingRoomGifts(db)
}
