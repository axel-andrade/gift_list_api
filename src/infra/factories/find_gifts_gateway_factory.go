package factories

import (
	repositories_impl "go_gift_list_api/src/infra/impl/repositories"
)

type FindGiftsGatewayFactory struct {
	*repositories_impl.GiftRepositoryImpl
}

func BuildFindGiftsGatewayFactory() *FindGiftsGatewayFactory {
	return &FindGiftsGatewayFactory{GiftRepositoryImpl: repositories_impl.BuildGiftRepositoryImpl()}
}
