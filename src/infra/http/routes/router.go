package routes

import (
	composes "go_gift_list_api/src/infra/composes"
	common_validators "go_gift_list_api/src/infra/http/middlewares/validators/common"
	gifts_validators "go_gift_list_api/src/infra/http/middlewares/validators/gifts"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
		main.GET("healthcheck", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "OK"})
		})
	}

	v1 := router.Group("/api/v1")
	{

		categories := v1.Group("categories")
		{
			categories.GET("", common_validators.PaginationValidator(), composes.FindCategoriesCompose)
		}

		gifts := v1.Group("gifts")
		{
			gifts.GET("", common_validators.PaginationValidator(), gifts_validators.FindGiftsValidator(), composes.FindGiftsCompose)
			gifts.GET("/marked", common_validators.PaginationValidator(), composes.FindMarkedGiftsCompose)
			gifts.POST("/mark", composes.MarkGiftCompose)
		}
	}

	return router
}
