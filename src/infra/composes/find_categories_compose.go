package composes

import (
	"go_gift_list_api/src/adapters/controllers"
	"go_gift_list_api/src/adapters/presenters"
	"go_gift_list_api/src/entities"
	"go_gift_list_api/src/infra/factories"
	interactor "go_gift_list_api/src/usecases/find_categories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindCategoriesCompose(c *gin.Context) {
	gateway := factories.BuildFindCategoriesGatewayFactory()
	ptr := presenters.FindCategoriesPresenter{}
	interactor := interactor.BuildFindCategoriesInteractor(gateway)
	ctrl := controllers.FindCategoriesController{Interactor: *interactor, Presenter: ptr}

	var paginationOptions entities.PaginationOptions
	paginationOptions.Limit, _ = strconv.Atoi(c.Query("limit"))
	paginationOptions.Page, _ = strconv.Atoi(c.Query("page"))
	paginationOptions.Sort = c.Query("order")

	output := ctrl.Run(paginationOptions)
	c.JSON(http.StatusOK, output)
}
