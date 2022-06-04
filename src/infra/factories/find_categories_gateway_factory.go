package factories

import (
	repositories_impl "go_gift_list_api/src/infra/impl/repositories"
)

type FindCategoriesGatewayFactory struct {
	*repositories_impl.CategoryRepositoryImpl
}

func BuildFindCategoriesGatewayFactory() *FindCategoriesGatewayFactory {
	return &FindCategoriesGatewayFactory{CategoryRepositoryImpl: repositories_impl.BuildCategoryRepositoryImpl()}
}
