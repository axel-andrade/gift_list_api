package presenters

import (
	common_adapters "go_gift_list_api/src/adapters/common"
	common_ptr "go_gift_list_api/src/adapters/presenters/common"
	"go_gift_list_api/src/entities"
	ERROR "go_gift_list_api/src/shared/errors"
	find_gifts_interactor "go_gift_list_api/src/usecases/find_gifts"

	"net/http"
)

type FindGiftsPresenter struct {
	giftPtr       common_ptr.GiftPresenter
	paginationPtr common_ptr.PaginationPresenter
}

func (p *FindGiftsPresenter) Show(result *find_gifts_interactor.FindGiftsOutputDTO, paginationOptions *entities.PaginationOptions, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatErrOutput(err)
	}

	return p.formatSuccessOutput(result, paginationOptions)
}

func (p *FindGiftsPresenter) formatSuccessOutput(result *find_gifts_interactor.FindGiftsOutputDTO, paginationOptions *entities.PaginationOptions) common_adapters.OutputPort {
	docsFormatted := p.giftPtr.FormatList(result.Gifts)
	data := p.paginationPtr.Format(docsFormatted, paginationOptions, result.TotalGifts)

	return common_adapters.OutputPort{
		StatusCode: http.StatusOK,
		Data:       data,
	}
}

func (p *FindGiftsPresenter) formatErrOutput(err error) common_adapters.OutputPort {
	errorMsg := err.Error()

	switch errorMsg {
	case ERROR.CATEGORY_NOT_FOUND:
		return common_adapters.OutputPort{
			StatusCode: http.StatusNotFound,
			Error:      errorMsg,
		}
	default:
		return common_adapters.OutputPort{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}
	}
}
