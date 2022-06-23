package mappers

import (
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/database/models"
)

type GiftMapper struct {
	BaseMapper
	CategoryMapper
}

func (m *GiftMapper) ToDomain(model models.Gift) *entities.Gift {
	var availableToDomain bool

	if model.Available == 0 {
		availableToDomain = false
	} else {
		availableToDomain = true
	}

	entity := &entities.Gift{
		Base:       *m.BaseMapper.toDomain(model.Base),
		Name:       model.Name,
		CategoryID: model.CategoryID,
		Available:  availableToDomain,
		Image:      model.Image,
		Quantity:   model.Quantity,
		PriceGrade: model.PriceGrade,
	}

	if model.Category.ID != 0 {
		entity.Category = *m.CategoryMapper.ToDomain(model.Category)
	}

	return entity
}

func (m *GiftMapper) ToPersistence(entity entities.Gift) *models.Gift {

	var availableToPersistent int64

	if entity.Available {
		availableToPersistent = 1
	} else {
		availableToPersistent = 0
	}

	return &models.Gift{
		Base:       *m.BaseMapper.toPersistence(entity.Base),
		Name:       entity.Name,
		CategoryID: entity.CategoryID,
		Available:  availableToPersistent,
		Image:      entity.Image,
		Quantity:   entity.Quantity,
		PriceGrade: entity.PriceGrade,
	}
}
