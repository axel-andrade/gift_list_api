package controllers

import (
	common_adapters "go_gift_list_api/src/adapters/common"
	"go_gift_list_api/src/adapters/presenters"
	interactor "go_gift_list_api/src/usecases/find_categories"
)

type FindCategoriesController struct {
	Interactor interactor.FindCategoriesInteractor
	Presenter  presenters.FindCategoriesPresenter
}

func (ctrl *FindCategoriesController) Run(input interactor.FindUserInputDTO) common_adapters.OutputPort {
	result, err := ctrl.Interactor.Execute(input)
	out := ctrl.Presenter.Show(result, err)
	return out
}
