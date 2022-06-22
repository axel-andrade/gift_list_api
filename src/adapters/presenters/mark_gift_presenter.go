package presenters

import (
	common_adapters "go_gift_list_api/src/adapters/common"
	ERROR "go_gift_list_api/src/shared/errors"

	"net/http"
)

type MarkGiftPresenter struct{}

func (p *MarkGiftPresenter) Show(err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatErrOutput(err)
	}

	return common_adapters.OutputPort{
		StatusCode: http.StatusCreated,
		Data:       struct{ msg string }{"successfully marked gift"},
	}
}

func (p *MarkGiftPresenter) formatErrOutput(err error) common_adapters.OutputPort {
	var statusCode int

	errorMsg := err.Error()

	switch errorMsg {
	case ERROR.CATEGORY_NOT_FOUND:
	case ERROR.GIFT_NOT_FOUND:
		statusCode = http.StatusNotFound
	case ERROR.GIFT_NOT_AVAILABLE:
	case ERROR.GIFT_QUANTITY_NOT_AVAILABLE:
		statusCode = http.StatusConflict
	default:
		statusCode = http.StatusInternalServerError
	}

	return common_adapters.OutputPort{
		StatusCode: statusCode,
		Error:      errorMsg,
	}
}
