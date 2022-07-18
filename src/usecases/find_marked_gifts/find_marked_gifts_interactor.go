package find_marked_gifts_interactor

type FindMarkedGiftsInteractor struct {
	Gateway FindMarkedGiftsGateway
}

func BuildFindMarkedGiftsInteractor(g FindMarkedGiftsGateway) *FindMarkedGiftsInteractor {
	return &FindMarkedGiftsInteractor{Gateway: g}
}

func (bs *FindMarkedGiftsInteractor) Execute(input FindMarkedGiftsInputDTO) (*FindMarkedGiftsOutputDTO, error) {
	markedGifts, total, err := bs.Gateway.FindMarkedGiftsPaginate(input.PaginationOptions)
	if err != nil {
		return nil, err
	}

	return &FindMarkedGiftsOutputDTO{markedGifts, total}, nil
}
