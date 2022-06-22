package factories

import (
	repositories_impl "go_gift_list_api/src/infra/impl/repositories"
)

type MarkGiftGatewayFactory struct {
	*repositories_impl.BaseRepositoryImpl
	*repositories_impl.CategoryRepositoryImpl
	*repositories_impl.GiftRepositoryImpl
	*repositories_impl.MarkedGiftRepositoryImpl
}

func BuildMarkGiftGatewayFactory() *MarkGiftGatewayFactory {
	return &MarkGiftGatewayFactory{
		BaseRepositoryImpl:       repositories_impl.BuildBaseRepoImpl(),
		CategoryRepositoryImpl:   repositories_impl.BuildCategoryRepositoryImpl(),
		GiftRepositoryImpl:       repositories_impl.BuildGiftRepositoryImpl(),
		MarkedGiftRepositoryImpl: repositories_impl.BuildMarkedGiftRepositoryImpl(),
	}
}
