package find_categories

import (
	"go_gift_list_api/src/entities"
)

type FindCategoriesGateway interface {
	FindCategoriesPaginate(pagination *entities.PaginationOptions) (*entities.PaginateResult, error)
}

type FindUserInputDTO = entities.PaginationOptions

type FindCategoriesOutputDTO = entities.PaginateResult
