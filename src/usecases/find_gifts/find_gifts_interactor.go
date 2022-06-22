package find_gifts_interactor

import (
	"errors"
	ERROR "go_gift_list_api/src/shared/errors"
)

type FindGiftsInteractor struct {
	Gateway FindGiftsGateway
}

func BuildFindGiftsInteractor(g FindGiftsGateway) *FindGiftsInteractor {
	return &FindGiftsInteractor{Gateway: g}
}

func (bs *FindGiftsInteractor) Execute(input *FindGiftsInputDTO) (*FindGiftsOutputDTO, error) {
	categoryID, paginationOptions := input.CategoryID, input.PaginationOptions

	if categoryID > 0 {
		categoryExists := bs.Gateway.CheckExistsCategory(categoryID)

		if !categoryExists {
			return nil, errors.New(ERROR.CATEGORY_NOT_FOUND)
		}
	}

	gifts, totalGifts, err := bs.Gateway.FindGiftsPaginate(categoryID, paginationOptions)
	if err != nil {
		return nil, err
	}

	return &FindGiftsOutputDTO{gifts, totalGifts}, nil
}
