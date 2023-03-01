package midleware

import (
	"net/http"
	"strings"

	"test_sat/pkg/auth"

	"test_sat/model"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
		err := auth.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.View{
				Message:      "Error",
				ErrorMessage: err.Error(),
			})
			c.Abort()
		}
		c.Next()
	}
}
