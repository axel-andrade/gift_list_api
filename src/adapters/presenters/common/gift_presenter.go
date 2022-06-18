package common_ptr

import (
	"go_gift_list_api/src/entities"
	"time"
)

type GiftFormatted struct {
	ID            entities.UniqueEntityID `json:"id" bson:"id"`
	Name          string                  `json:"name" bson:"name"`
	Category      CategoryFormatted       `json:"category" bson:"category"`
	Available     bool                    `json:"available" bson:"available"`
	ImageFilename string                  `json:"image_filename" bson:"image_filename"`
	Quantity      int64                   `json:"quantity" bson:"quantity"`
	CreatedAt     time.Time               `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time               `json:"updated_at" bson:"updated_at"`
}

type GiftPresenter struct {
	categoryPtr CategoryPresenter
}

func (ptr *GiftPresenter) Format(gift entities.Gift) GiftFormatted {
	return GiftFormatted{
		ID:            gift.ID,
		Name:          gift.Name,
		Category:      ptr.categoryPtr.Format(gift.Category),
		Available:     gift.Available,
		ImageFilename: gift.ImageFilename,
		Quantity:      gift.Quantity,
		CreatedAt:     gift.CreatedAt,
		UpdatedAt:     gift.UpdatedAt,
	}
}

func (ptr *GiftPresenter) FormatList(gifts []entities.Gift) []GiftFormatted {
	var giftsFormatted []GiftFormatted

	for _, gift := range gifts {
		giftsFormatted = append(giftsFormatted, ptr.Format(gift))
	}

	return giftsFormatted
}
