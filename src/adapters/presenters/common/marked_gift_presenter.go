package common_ptr

import (
	"go_gift_list_api/src/entities"
	"time"
)

type MarkedGiftFormatted struct {
	ID         entities.UniqueEntityID `json:"id" bson:"id"`
	PersonName string                  `json:"person_name" bson:"person_name"`
	GiftName   string                  `json:"gift_name" bson:"gift_name"`
	CreatedAt  time.Time               `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time               `json:"updated_at" bson:"updated_at"`
}

type MarkedGiftPresenter struct{}

func (ptr *MarkedGiftPresenter) Format(MarkedGift entities.MarkedGift) MarkedGiftFormatted {
	return MarkedGiftFormatted{
		ID:         MarkedGift.ID,
		PersonName: MarkedGift.PersonName,
		GiftName:   MarkedGift.Gift.Name,
		CreatedAt:  MarkedGift.CreatedAt,
		UpdatedAt:  MarkedGift.UpdatedAt,
	}
}

func (ptr *MarkedGiftPresenter) FormatList(gifts []entities.MarkedGift) []MarkedGiftFormatted {
	var markedGiftsFormatted []MarkedGiftFormatted = make([]MarkedGiftFormatted, 0)

	for _, gift := range gifts {
		markedGiftsFormatted = append(markedGiftsFormatted, ptr.Format(gift))
	}

	return markedGiftsFormatted
}
