package find_categories_interactor

import (
	"go_gift_list_api/src/entities"
)

type FindCategoriesGateway interface {
	FindCategoriesPaginate(pagination *entities.PaginationOptions) ([]entities.Category, int64, error)
}

type FindCategoriesInputDTO struct {
	PaginationOptions *entities.PaginationOptions
}

type FindCategoriesOutputDTO struct {
	Categories      []entities.Category
	TotalCategories int64
}
