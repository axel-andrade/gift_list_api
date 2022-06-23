package factories

import (
	repositories_impl "go_gift_list_api/src/infra/impl/repositories"
)

type FindGiftsGatewayFactory struct {
	*repositories_impl.GiftRepositoryImpl
}

func BuildFindGiftsGatewayFactory() *FindGiftsGatewayFactory {
	baseRepositoryImpl := repositories_impl.BuildBaseRepoImpl()
	categoryRepositoryImpl := repositories_impl.BuildCategoryRepositoryImpl(baseRepositoryImpl)
	giftRepositoryImpl := repositories_impl.BuildGiftRepositoryImpl(baseRepositoryImpl, categoryRepositoryImpl)

	return &FindGiftsGatewayFactory{GiftRepositoryImpl: giftRepositoryImpl}
}
