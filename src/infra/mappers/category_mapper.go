package mappers

import (
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/database/models"
)

type CategoryMapper struct {
	BaseMapper
}

func (m *CategoryMapper) ToDomain(model models.Category) *entities.Category {
	return &entities.Category{
		Base: *m.BaseMapper.toDomain(model.Base),
		Name: model.Name,
	}
}

func (m *CategoryMapper) ToPersistence(entity entities.Category) *models.Category {
	return &models.Category{
		Base: *m.BaseMapper.toPersistence(entity.Base),
		Name: entity.Name,
	}
}
