package find_categories

type FindCategoriesInteractor struct {
	Gateway FindCategoriesGateway
}

func BuildFindCategoriesInteractor(g FindCategoriesGateway) *FindCategoriesInteractor {
	return &FindCategoriesInteractor{Gateway: g}
}

func (bs *FindCategoriesInteractor) Execute(input FindUserInputDTO) (*FindCategoriesOutputDTO, error) {
	data, err := bs.Gateway.FindCategoriesPaginate(&input)
	if err != nil {
		return nil, err
	}

	return data, nil
}
