package common_ptr

import (
	"go_gift_list_api/src/entities"
	"time"
)

type CategoryFormatted struct {
	ID        entities.UniqueEntityID `json:"id" bson:"id"`
	Name      string                  `json:"name" bson:"name"`
	CreatedAt time.Time               `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time               `json:"updated_at" bson:"updated_at"`
}

type CategoryPresenter struct{}

func (ptr *CategoryPresenter) Format(category entities.Category) CategoryFormatted {
	return CategoryFormatted{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func (ptr *CategoryPresenter) FormatList(gifts []entities.Category) []CategoryFormatted {
	var categoriesFormatted []CategoryFormatted

	for _, gift := range gifts {
		categoriesFormatted = append(categoriesFormatted, ptr.Format(gift))
	}

	return categoriesFormatted
}
