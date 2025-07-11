package repositories_impl

import (
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/database/models"
	"go_gift_list_api/src/infra/mappers"
	"strconv"
)

type GiftRepositoryImpl struct {
	*BaseRepositoryImpl
	*CategoryRepositoryImpl
	GiftMapper mappers.GiftMapper
}

func BuildGiftRepositoryImpl(baseRepositoryImpl *BaseRepositoryImpl, categoryRepositoryImpl *CategoryRepositoryImpl) *GiftRepositoryImpl {
	return &GiftRepositoryImpl{BaseRepositoryImpl: baseRepositoryImpl, CategoryRepositoryImpl: categoryRepositoryImpl}
}

func (r *GiftRepositoryImpl) FindGiftByID(id entities.UniqueEntityID) (*entities.Gift, error) {
	var gift entities.Gift
	err := r.Db.First(&gift, "id = ?", id).Error

	if err != nil || gift.ID == 0 {
		return nil, err
	}

	return &gift, nil
}

func (r *GiftRepositoryImpl) FindGiftsPaginate(categoryID entities.UniqueEntityID, pagination *entities.PaginationOptions) ([]entities.Gift, int64, error) {
	var count int64
	var q string = "available = 1"

	if categoryID > 0 {
		q += " AND category_id = " + strconv.Itoa(int(categoryID))
	}

	if pagination.Search != "" {
		q += " AND name LIKE '%" + pagination.Search + "%'"
	}

	r.Db.Model(&models.Gift{}).Where(q).Count(&count)

	offset := pagination.GetOffset()
	limit := pagination.Limit
	sort := "price_grade desc"

	var gifts []entities.Gift

	r.Db.Preload("Category").Offset(offset).Limit(limit).Order(sort).Where(q).Find(&gifts)

	return gifts, count, nil
}

func (r *GiftRepositoryImpl) UpdateGift(gift *entities.Gift) (*entities.Gift, error) {
	giftModel := r.GiftMapper.ToPersistence(*gift)

	q := r.getQueryOrTx()

	result := q.Model(&models.Gift{}).Where("id = ?", gift.ID).Updates(&map[string]interface{}{"available": giftModel.Available, "quantity": giftModel.Quantity})

	if result.Error != nil {
		return nil, result.Error
	}

	return r.GiftMapper.ToDomain(*giftModel), nil
}
