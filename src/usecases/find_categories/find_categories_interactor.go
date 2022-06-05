package find_categories

import "go_gift_list_api/src/entities"

type FindCategoriesInteractor struct {
	Gateway FindCategoriesGateway
}

func BuildFindCategoriesInteractor(g FindCategoriesGateway) *FindCategoriesInteractor {
	return &FindCategoriesInteractor{Gateway: g}
}

func (bs *FindCategoriesInteractor) Execute(input FindUserInputDTO) (*FindCategoriesOutputDTO, error) {
	// data, err := bs.Gateway.FindCategoriesPaginate(&input)
	// if err != nil {
	// 	return nil, err
	// }

	// return data, nil
	return &entities.PaginateResult{
		Docs:        nil,
		TotalDocs:   0,
		TotalPages:  0,
		Limit:       0,
		Page:        0,
		HasPrevPage: true,
		HasNextPage: true,
		PrevPage:    0,
		NextPage:    0,
	}, nil
}
