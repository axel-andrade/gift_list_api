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

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	sort := c.Query("sort")

	paginationOptions, _ := entities.BuildPaginationOptions(limit, page, sort)

	input := find_gifts_interactor.FindsGiftsInputDTO{PaginationOptions: paginationOptions}

	output := ctrl.Run(input)
	c.JSON(http.StatusOK, output)
}
