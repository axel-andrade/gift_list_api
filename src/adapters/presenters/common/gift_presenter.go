package common_ptr

import (
	"go_gift_list_api/src/entities"
	"time"
)

type GiftFormatted struct {
	ID         entities.UniqueEntityID `json:"id" bson:"id"`
	Name       string                  `json:"name" bson:"name"`
	Category   CategoryFormatted       `json:"category" bson:"category"`
	Available  bool                    `json:"available" bson:"available"`
	Image      string                  `json:"image" bson:"image"`
	Quantity   int64                   `json:"quantity" bson:"quantity"`
	PriceGrade int64                   `json:"price_grade" bson:"price_grade"`
	CreatedAt  time.Time               `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time               `json:"updated_at" bson:"updated_at"`
}

type GiftPresenter struct {
	categoryPtr CategoryPresenter
}

func (ptr *GiftPresenter) Format(gift entities.Gift) GiftFormatted {
	return GiftFormatted{
		ID:         gift.ID,
		Name:       gift.Name,
		Category:   ptr.categoryPtr.Format(gift.Category),
		Available:  gift.Available,
		Image:      gift.Image,
		Quantity:   gift.Quantity,
		PriceGrade: gift.PriceGrade,
		CreatedAt:  gift.CreatedAt,
		UpdatedAt:  gift.UpdatedAt,
	}
}

func (ptr *GiftPresenter) FormatList(gifts []entities.Gift) []GiftFormatted {
	var giftsFormatted []GiftFormatted = make([]GiftFormatted, 0)

	for _, gift := range gifts {
		giftsFormatted = append(giftsFormatted, ptr.Format(gift))
	}

	return giftsFormatted
}
