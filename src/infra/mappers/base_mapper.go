package mappers

import (
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/database/models"
)

type BaseMapper struct{}

func (m *BaseMapper) toDomain(model models.Base) *entities.Base {
	return &entities.Base{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (m *BaseMapper) toPersistence(entity entities.Base) *models.Base {
	return &models.Base{
		ID:        entity.ID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
