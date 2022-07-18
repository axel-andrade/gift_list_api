package controllers

import (
	common_adapters "go_gift_list_api/src/adapters/common"
	"go_gift_list_api/src/adapters/presenters"
	interactor "go_gift_list_api/src/usecases/find_marked_gifts"
)

type FindMarkedGiftsController struct {
	Interactor interactor.FindMarkedGiftsInteractor
	Presenter  presenters.FindMarkedGiftsPresenter
}

func (ctrl *FindMarkedGiftsController) Run(input interactor.FindMarkedGiftsInputDTO) common_adapters.OutputPort {
	result, err := ctrl.Interactor.Execute(input)
	return ctrl.Presenter.Show(result, input.PaginationOptions, err)
}
