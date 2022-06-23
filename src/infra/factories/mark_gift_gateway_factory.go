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
	baseRepositoryImpl := repositories_impl.BuildBaseRepoImpl()
	categoryRepositoryImpl := repositories_impl.BuildCategoryRepositoryImpl(baseRepositoryImpl)
	giftRepositoryImpl := repositories_impl.BuildGiftRepositoryImpl(baseRepositoryImpl, categoryRepositoryImpl)
	markedGiftRepositoryImpl := repositories_impl.BuildMarkedGiftRepositoryImpl(baseRepositoryImpl)

	return &MarkGiftGatewayFactory{
		BaseRepositoryImpl:       baseRepositoryImpl,
		CategoryRepositoryImpl:   categoryRepositoryImpl,
		GiftRepositoryImpl:       giftRepositoryImpl,
		MarkedGiftRepositoryImpl: markedGiftRepositoryImpl,
	}
}
