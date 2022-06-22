package composes

import (
	"go_gift_list_api/src/adapters/controllers"
	"go_gift_list_api/src/adapters/presenters"
	"go_gift_list_api/src/infra/factories"
	mark_gift_interactor "go_gift_list_api/src/usecases/mark_gift"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MarkGiftCompose(c *gin.Context) {
	var input mark_gift_interactor.MarkGiftInputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	gateway := factories.BuildMarkGiftGatewayFactory()
	ptr := presenters.MarkGiftPresenter{}
	interactor := mark_gift_interactor.BuildMarkGiftInteractor(gateway)
	ctrl := controllers.MarkGiftController{Interactor: *interactor, Presenter: ptr}

	output := ctrl.Run(&input)
	c.JSON(http.StatusOK, output)
}
