package repositories

import (
	"go_gift_list_api/src/entities"
)

type CategoryRepository interface {
	BaseRepository
	CreateCategory(category *entities.Category) (*entities.Category, error)
	UpdateCategory(category *entities.Category) error
	FindCategoryByID(id entities.UniqueEntityID) (*entities.Category, error)
	FindCategoriesPaginate(pagination *entities.PaginationOptions) ([]entities.Category, int64, error)
}
