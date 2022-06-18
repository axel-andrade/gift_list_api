package presenters

import (
	common_adapters "go_gift_list_api/src/adapters/common"
	common_ptr "go_gift_list_api/src/adapters/presenters/common"
	"go_gift_list_api/src/entities"
	interactor "go_gift_list_api/src/usecases/find_categories"
	"net/http"
)

type FindCategoriesPresenter struct {
	categoryPtr   common_ptr.CategoryPresenter
	paginationPtr common_ptr.PaginationPresenter
}

func (p *FindCategoriesPresenter) Show(result *interactor.FindCategoriesOutputDTO, paginationOptions entities.PaginationOptions, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatErrOutput(err)
	}

	return p.formatSuccessOutput(result, paginationOptions)
}

func (p *FindCategoriesPresenter) formatSuccessOutput(result *interactor.FindCategoriesOutputDTO, paginationOptions entities.PaginationOptions) common_adapters.OutputPort {
	docs := p.categoryPtr.FormatList(result.Categories)
	data := p.paginationPtr.Format(docs, paginationOptions, result.TotalCategories)

	return common_adapters.OutputPort{
		StatusCode: http.StatusOK,
		Data:       data,
	}
}

func (p *FindCategoriesPresenter) formatErrOutput(err error) common_adapters.OutputPort {
	return common_adapters.OutputPort{
		StatusCode: http.StatusBadRequest,
		Error:      err.Error(),
	}
}
