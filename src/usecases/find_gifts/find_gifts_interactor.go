package find_gifts_interactor

type FindGiftsInteractor struct {
	Gateway FindGiftsGateway
}

func BuildFindGiftsInteractor(g FindGiftsGateway) *FindGiftsInteractor {
	return &FindGiftsInteractor{Gateway: g}
}

func (bs *FindGiftsInteractor) Execute(input *FindsGiftsInputDTO) (*FindsGiftsOutputDTO, error) {
	categoryID, paginationOptions := input.CategoryID, input.PaginationOptions

	gifts, totalGifts, err := bs.Gateway.FindGiftsPaginate(categoryID, paginationOptions)
	if err != nil {
		return nil, err
	}

	return &FindsGiftsOutputDTO{gifts, totalGifts}, nil
}
