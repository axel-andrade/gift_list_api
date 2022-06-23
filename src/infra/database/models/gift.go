package models

type Gift struct {
	Base
	Name       string   `gorm:"type:varchar(255);not null" json:"name"`
	CategoryID int64    `json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Available  int64    `gorm:"type:tinyint(1);default:1" json:"available"`
	Image      string   `gorm:"type:varchar(65000);not null" json:"image"`
	Quantity   int64    `gorm:"type:integer;default:1" json:"quantity"`
	PriceGrade int64    `gorm:"type:integer;default:1" json:"price_grade"`
}
