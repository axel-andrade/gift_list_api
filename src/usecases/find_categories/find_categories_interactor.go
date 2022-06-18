package find_categories_interactor

type FindCategoriesInteractor struct {
	Gateway FindCategoriesGateway
}

func BuildFindCategoriesInteractor(g FindCategoriesGateway) *FindCategoriesInteractor {
	return &FindCategoriesInteractor{Gateway: g}
}

func (bs *FindCategoriesInteractor) Execute(input FindCategoriesInputDTO) (*FindCategoriesOutputDTO, error) {
	categories, totalCategories, err := bs.Gateway.FindCategoriesPaginate(input.PaginationOptions)
	if err != nil {
		return nil, err
	}

	return &FindCategoriesOutputDTO{categories, totalCategories}, nil
}
