package repositories_impl

import (
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/database/models"
	"go_gift_list_api/src/infra/mappers"
)

type CategoryRepositoryImpl struct {
	*BaseRepositoryImpl
	CategoryMapper mappers.CategoryMapper
}

func BuildCategoryRepositoryImpl(baseRepositoryImpl *BaseRepositoryImpl) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{BaseRepositoryImpl: baseRepositoryImpl}
}

func (r *CategoryRepositoryImpl) CheckExistsCategory(categoryID entities.UniqueEntityID) bool {
	var count int64

	r.Db.Model(&models.Category{}).Where("id = ?", categoryID).Count(&count)

	return count > 0
}

func (r *CategoryRepositoryImpl) CreateCategory(category *entities.Category) (*entities.Category, error) {

	model := r.CategoryMapper.ToPersistence(*category)

	q := r.getQueryOrTx()

	err := q.Create(model).Error

	if err != nil {
		return nil, err
	}

	return r.CategoryMapper.ToDomain(*model), nil
}

func (r *CategoryRepositoryImpl) FindCategoriesPaginate(pagination *entities.PaginationOptions) ([]entities.Category, int64, error) {
	var count int64
	r.Db.Model(&models.Category{}).Count(&count)

	offset := pagination.GetOffset()
	limit := pagination.Limit
	sort := pagination.Sort

	var categories []entities.Category
	r.Db.Offset(offset).Limit(limit).Order(sort).Find(&categories)

	return categories, count, nil
}

func (r *CategoryRepositoryImpl) FindCategoryByID(id entities.UniqueEntityID) (*entities.Category, error) {
	var category entities.Category
	err := r.Db.First(&category, "id = ?", id).Error

	if err != nil || category.ID > 0 {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepositoryImpl) UpdateCategory(category *entities.Category) error {
	err := r.Db.Save(category).Error
	return err
}
