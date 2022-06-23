package models

type MarkedGift struct {
	Base
	PersonName string `gorm:"type:varchar(255);not null" json:"person_name"`
	GiftID     int64  `json:"gift_id"`
	Gift       Gift   `gorm:"foreignKey:GiftID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity   int64  `gorm:"type:integer;default:1" json:"quantity"`
}
