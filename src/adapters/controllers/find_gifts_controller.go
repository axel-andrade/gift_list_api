package controllers

import (
	common_adapters "go_gift_list_api/src/adapters/common"
	"go_gift_list_api/src/adapters/presenters"
	interactor "go_gift_list_api/src/usecases/find_gifts"
)

type FindGiftsController struct {
	Interactor interactor.FindGiftsInteractor
	Presenter  presenters.FindGiftsPresenter
}

func (ctrl *FindGiftsController) Run(input *interactor.FindGiftsInputDTO) common_adapters.OutputPort {
	result, err := ctrl.Interactor.Execute(input)
	return ctrl.Presenter.Show(result, input.PaginationOptions, err)
}
