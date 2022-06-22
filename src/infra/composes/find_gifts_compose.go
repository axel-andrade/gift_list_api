package composes

import (
	"go_gift_list_api/src/adapters/controllers"
	"go_gift_list_api/src/adapters/presenters"
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/factories"
	find_gifts_interactor "go_gift_list_api/src/usecases/find_gifts"

	"github.com/gin-gonic/gin"
)

func FindGiftsCompose(c *gin.Context) {
	gateway := factories.BuildFindGiftsGatewayFactory()
	ptr := presenters.FindGiftsPresenter{}
	interactor := find_gifts_interactor.BuildFindGiftsInteractor(gateway)
	ctrl := controllers.FindGiftsController{Interactor: *interactor, Presenter: ptr}

	categoryID := int64(c.MustGet("categoryID").(int))
	paginationOptions := c.MustGet("paginationOptions").(*entities.PaginationOptions)

	input := find_gifts_interactor.FindGiftsInputDTO{CategoryID: categoryID, PaginationOptions: paginationOptions}

	output := ctrl.Run(&input)

	c.JSON(output.StatusCode, output)
}
