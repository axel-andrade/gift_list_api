package mappers

import (
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/database/models"
)

type MarkedGiftMapper struct {
	BaseMapper
	CategoryMapper
	GiftMapper
}

func (m *MarkedGiftMapper) ToDomain(model models.MarkedGift) *entities.MarkedGift {
	entity := &entities.MarkedGift{
		Base:       *m.BaseMapper.toDomain(model.Base),
		PersonName: model.PersonName,
		GiftID:     model.GiftID,
		Quantity:   model.Quantity,
	}

	if model.Gift.ID != 0 {
		entity.Gift = *m.GiftMapper.ToDomain(model.Gift)
	}

	return entity
}

func (m *MarkedGiftMapper) ToPersistence(entity entities.MarkedGift) *models.MarkedGift {
	return &models.MarkedGift{
		Base:       *m.BaseMapper.toPersistence(entity.Base),
		PersonName: entity.PersonName,
		GiftID:     entity.GiftID,
		Quantity:   entity.Quantity,
	}
}
