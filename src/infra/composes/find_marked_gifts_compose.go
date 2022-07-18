package composes

import (
	"go_gift_list_api/src/adapters/controllers"
	"go_gift_list_api/src/adapters/presenters"
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/factories"
	find_marked_gifts_interactor "go_gift_list_api/src/usecases/find_marked_gifts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindMarkedGiftsCompose(c *gin.Context) {
	gateway := factories.BuildMarkGiftGatewayFactory()
	ptr := presenters.FindMarkedGiftsPresenter{}
	interactor := find_marked_gifts_interactor.BuildFindMarkedGiftsInteractor(gateway)
	ctrl := controllers.FindMarkedGiftsController{Interactor: *interactor, Presenter: ptr}

	paginationOptions := c.MustGet("paginationOptions").(*entities.PaginationOptions)

	input := find_marked_gifts_interactor.FindMarkedGiftsInputDTO{PaginationOptions: paginationOptions}

	output := ctrl.Run(input)
	c.JSON(http.StatusOK, output)
}
