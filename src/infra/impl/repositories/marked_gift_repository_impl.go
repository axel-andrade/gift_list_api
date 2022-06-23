package repositories_impl

import (
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/mappers"
)

type MarkedGiftRepositoryImpl struct {
	*BaseRepositoryImpl
	MarkedGiftMapper mappers.MarkedGiftMapper
}

func BuildMarkedGiftRepositoryImpl(baseRepositoryImpl *BaseRepositoryImpl) *MarkedGiftRepositoryImpl {
	return &MarkedGiftRepositoryImpl{BaseRepositoryImpl: baseRepositoryImpl}
}

func (r *MarkedGiftRepositoryImpl) CreateMarkedGift(markedGift *entities.MarkedGift) (*entities.MarkedGift, error) {

	model := r.MarkedGiftMapper.ToPersistence(*markedGift)

	q := r.getQueryOrTx()

	err := q.Create(model).Error

	if err != nil {
		return nil, err
	}

	return r.MarkedGiftMapper.ToDomain(*model), nil
}
