package composes

import (
	"go_gift_list_api/src/adapters/controllers"
	"go_gift_list_api/src/adapters/presenters"
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/factories"
	find_gifts_interactor "go_gift_list_api/src/usecases/find_gifts"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindGiftsCompose(c *gin.Context) {
	gateway := factories.BuildFindGiftsGatewayFactory()
	ptr := presenters.FindGiftsPresenter{}
	interactor := find_gifts_interactor.BuildFindGiftsInteractor(gateway)
	ctrl := controllers.FindGiftsController{Interactor: *interactor, Presenter: ptr}

	var paginationOptions entities.PaginationOptions
	paginationOptions.Limit, _ = strconv.Atoi(c.Query("limit"))
	paginationOptions.Page, _ = strconv.Atoi(c.Query("page"))
	paginationOptions.Sort = c.Query("order")

	input := find_gifts_interactor.FindsGiftsInputDTO{PaginationOptions: paginationOptions}

	output := ctrl.Run(input)
	c.JSON(http.StatusOK, output)
}
