package presenters

import (
	common_adapters "go_gift_list_api/src/adapters/common"
	common_ptr "go_gift_list_api/src/adapters/presenters/common"
	"go_gift_list_api/src/entities"
	interactor "go_gift_list_api/src/usecases/find_marked_gifts"
	"net/http"
)

type FindMarkedGiftsPresenter struct {
	markedGiftPtr common_ptr.MarkedGiftPresenter
	paginationPtr common_ptr.PaginationPresenter
}

func (p *FindMarkedGiftsPresenter) Show(result *interactor.FindMarkedGiftsOutputDTO, paginationOptions *entities.PaginationOptions, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatErrOutput(err)
	}

	return p.formatSuccessOutput(result, paginationOptions)
}

func (p *FindMarkedGiftsPresenter) formatSuccessOutput(result *interactor.FindMarkedGiftsOutputDTO, paginationOptions *entities.PaginationOptions) common_adapters.OutputPort {
	docs := p.markedGiftPtr.FormatList(result.MarkedGifts)
	data := p.paginationPtr.Format(docs, paginationOptions, result.TotalMarkedGifts)

	return common_adapters.OutputPort{
		StatusCode: http.StatusOK,
		Data:       data,
	}
}

func (p *FindMarkedGiftsPresenter) formatErrOutput(err error) common_adapters.OutputPort {
	return common_adapters.OutputPort{
		StatusCode: http.StatusBadRequest,
		Error:      err.Error(),
	}
}
