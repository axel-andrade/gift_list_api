package controllers

import (
	common_adapters "go_gift_list_api/src/adapters/common"
	"go_gift_list_api/src/adapters/presenters"
	interactor "go_gift_list_api/src/usecases/mark_gift"
)

type MarkGiftController struct {
	Interactor interactor.MarkGiftInteractor
	Presenter  presenters.MarkGiftPresenter
}

func (ctrl *MarkGiftController) Run(input *interactor.MarkGiftInputDTO) common_adapters.OutputPort {
	err := ctrl.Interactor.Execute(input)
	return ctrl.Presenter.Show(err)
}
