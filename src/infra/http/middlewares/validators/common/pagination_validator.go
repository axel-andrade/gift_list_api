package common_validators

import (
	"net/http"
	"strconv"
	"strings"

	"go_gift_list_api/src/entities"
	ERROR "go_gift_list_api/src/shared/errors"

	"github.com/gin-gonic/gin"
)

func PaginationValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var limit, page int
		var err error

		if c.Query("limit") != "" {
			if limit, err = strconv.Atoi(c.Query("limit")); err != nil {
				msg := strings.Replace(ERROR.FIELD_MUST_BE_INTEGER, "[field]", "limit", 1)

				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": msg,
				})

				c.Abort()
				return
			}

			if limit <= 0 {
				msg := strings.Replace(ERROR.FIELD_MUST_BE_GREATER_THAN_ZERO, "[field]", "limit", 1)

				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": msg,
				})

				c.Abort()
				return
			}
		}

		if c.Query("page") != "" {
			if page, err = strconv.Atoi(c.Query("page")); err != nil {
				msg := strings.Replace(ERROR.FIELD_MUST_BE_INTEGER, "[field]", "page", 1)

				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": msg,
				})

				c.Abort()
				return
			}

			if page <= 0 {
				msg := strings.Replace(ERROR.FIELD_MUST_BE_GREATER_THAN_ZERO, "[field]", "page", 1)

				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": msg,
				})

				c.Abort()
				return
			}
		}

		sort := c.Query("sort")
		search := c.Query("search")

		paginationOptions, _ := entities.BuildPaginationOptions(limit, page, sort, search)

		c.Set("paginationOptions", paginationOptions)
		c.Next()
	}
}
