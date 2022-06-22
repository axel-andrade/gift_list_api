package entities

type MarkedGift struct {
	Base
	PersonName string         `json:"person_name"`
	GiftID     UniqueEntityID `json:"gift_id"`
	Gift       Gift           `json:"gift"`
	CategoryID UniqueEntityID `json:"category_id"`
	Category   Category       `json:"category"`
	Quantity   int64          `gorm:"type:integer;default:1" json:"quantity"`
}
