package presenters

import (
	common_adapters "go_gift_list_api/src/adapters/common"
	interactor "go_gift_list_api/src/usecases/find_categories"
	"net/http"
)

type FindCategoriesPresenter struct{}

func (p *FindCategoriesPresenter) Show(result *interactor.FindCategoriesOutputDTO, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatErrOutput(err)
	}

	return p.formatSuccessOutput(result)
}

func (p *FindCategoriesPresenter) formatSuccessOutput(result *interactor.FindCategoriesOutputDTO) common_adapters.OutputPort {
	return common_adapters.OutputPort{
		StatusCode: http.StatusOK,
		Data:       result,
	}
}

func (p *FindCategoriesPresenter) formatErrOutput(err error) common_adapters.OutputPort {
	return common_adapters.OutputPort{
		StatusCode: http.StatusBadRequest,
		Error:      err.Error(),
	}
}
