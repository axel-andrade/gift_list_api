package repositories

import (
	"go_gift_list_api/src/entities"
)

type CategoryRepository interface {
	BaseRepository
	CreateCategory(user *entities.Category) (*entities.Category, error)
	UpdateCategory(user *entities.Category) error
	FindCategoryByID(id entities.UniqueEntityID) (*entities.Category, error)
	FindCategorysPaginate(pagination *entities.PaginationOptions) (*entities.PaginateResult, error)
}
