package repositories_impl

import (
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/database/models"
	"go_gift_list_api/src/infra/mappers"
)

type GiftRepositoryImpl struct {
	BaseRepositoryImpl
	CategoryRepositoryImpl
	GiftMapper mappers.GiftMapper
}

func BuildGiftRepositoryImpl() *GiftRepositoryImpl {
	return &GiftRepositoryImpl{BaseRepositoryImpl: *BuildBaseRepoImpl(), CategoryRepositoryImpl: *BuildCategoryRepositoryImpl()}
}

func (r *GiftRepositoryImpl) FindGiftByID(id entities.UniqueEntityID) (*entities.Gift, error) {
	var gift entities.Gift
	err := r.Db.First(&gift, "id = ?", id).Error

	if err != nil || gift.ID > 0 {
		return nil, err
	}

	return &gift, nil
}

func (r *GiftRepositoryImpl) FindGiftsPaginate(categoryID entities.UniqueEntityID, pagination *entities.PaginationOptions) ([]entities.Gift, int64, error) {
	var count int64
	var q string = "available = 1"

	if categoryID > 0 {
		q += " AND category_id = " + string(rune(categoryID))
	}

	if pagination.Search != "" {
		q += " AND name LIKE '%" + pagination.Search + "%'"
	}

	r.Db.Model(&models.Gift{}).Where(q).Count(&count)

	offset := pagination.GetOffset()
	limit := pagination.Limit
	sort := pagination.Sort

	var gifts []entities.Gift

	r.Db.Preload("Category").Offset(offset).Limit(limit).Order(sort).Where(q).Find(&gifts)

	return gifts, count, nil
}

func (r *GiftRepositoryImpl) UpdateGift(gift *entities.Gift) (*entities.Gift, error) {

	model := r.GiftMapper.ToPersistence(*gift)

	q := r.getQueryOrTx()

	if err := q.Where("id = ?", gift.ID).UpdateColumns(model).Error; err != nil {
		return nil, err
	}

	return r.GiftMapper.ToDomain(*model), nil
}
