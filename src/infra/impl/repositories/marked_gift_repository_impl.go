package repositories_impl

import (
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/database/models"
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

func (r *MarkedGiftRepositoryImpl) FindMarkedGiftsPaginate(pagination *entities.PaginationOptions) ([]entities.MarkedGift, int64, error) {
	var count int64
	r.Db.Model(&models.MarkedGift{}).Count(&count)

	offset := pagination.GetOffset()
	limit := pagination.Limit
	sort := pagination.Sort

	var markedGifts []entities.MarkedGift
	r.Db.Preload("Gift").Offset(offset).Limit(limit).Order(sort).Find(&markedGifts)

	return markedGifts, count, nil
}
