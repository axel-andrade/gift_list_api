package gifts_validators

import (
	"net/http"
	"strconv"
	"strings"

	ERROR "go_gift_list_api/src/shared/errors"

	"github.com/gin-gonic/gin"
)

func FindGiftsValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var categoryID int
		var err error

		if c.Query("category_id") != "" {
			if categoryID, err = strconv.Atoi(c.Query("category_id")); err != nil {
				msg := strings.Replace(ERROR.FIELD_MUST_BE_INTEGER, "[field]", "category_id", 1)

				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": msg,
				})

				c.Abort()
				return
			}

			if categoryID <= 0 {
				msg := strings.Replace(ERROR.FIELD_MUST_BE_GREATER_THAN_ZERO, "[field]", "category_id", 1)

				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"error": msg,
				})

				c.Abort()
				return
			}

			c.Set("categoryID", categoryID)
		} else {
			c.Set("categoryID", 0)
		}

		c.Next()
	}
}
