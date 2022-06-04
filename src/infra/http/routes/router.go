package routes

import (
	composes "go_gift_list_api/src/infra/composes"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("/")
	{
		main.GET("/healthcheck", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "OK"})
		})
	}

	v1 := router.Group("api/v1")
	{

		users := v1.Group("users")
		{
			users.GET("/", composes.FindUsersCompose)
		}
	}

	return router
}
