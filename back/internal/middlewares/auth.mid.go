package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/saegus/test-technique-romain-chenard/pkg/encrypt"
	"github.com/saegus/test-technique-romain-chenard/pkg/utils"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth_header, ok := c.Request.Header["Authorization"]
		if !ok || !strings.HasPrefix(auth_header[0], "Bearer") {
			c.JSON(http.StatusBadRequest, gin.H{"message": "token missing"})
			c.Abort()
			return
		}
		token := strings.Split(auth_header[0], " ")[1]
		// fmt.Println("got token : ", token)
		claim, err := encrypt.GetClaimsFromToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauhorized"})
			c.Abort()
			return
		}
		utils.PrettyDisplay(claim)

		c.Set("user_email", claim["Email"])
		c.Set("user_id", claim["ID"])
		c.Next()
	}
}
