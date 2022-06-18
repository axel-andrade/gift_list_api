package composes

import (
	"go_gift_list_api/src/adapters/controllers"
	"go_gift_list_api/src/adapters/presenters"
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/factories"
	find_categories_interactor "go_gift_list_api/src/usecases/find_categories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindCategoriesCompose(c *gin.Context) {
	gateway := factories.BuildFindCategoriesGatewayFactory()
	ptr := presenters.FindCategoriesPresenter{}
	interactor := find_categories_interactor.BuildFindCategoriesInteractor(gateway)
	ctrl := controllers.FindCategoriesController{Interactor: *interactor, Presenter: ptr}

	paginationOptions := c.MustGet("paginationOptions").(*entities.PaginationOptions)

	input := find_categories_interactor.FindCategoriesInputDTO{PaginationOptions: paginationOptions}

	output := ctrl.Run(input)
	c.JSON(http.StatusOK, output)
}
