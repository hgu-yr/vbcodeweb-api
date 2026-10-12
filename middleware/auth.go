package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"vbCodeWeb/common"
	"vbCodeWeb/utils"
)

// JWTAuth 校验 Authorization: Bearer <token>，通过后写入用户上下文。
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := ""
		if auth := c.GetHeader("Authorization"); strings.HasPrefix(auth, "Bearer ") {
			tokenStr = strings.TrimPrefix(auth, "Bearer ")
		} else {
			tokenStr = c.Query("token")
		}
		if tokenStr == "" {
			common.FailWith(common.CodeUnauth, "未登录").JSON(c)
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			common.FailWith(common.CodeUnauth, "登录已失效，请重新登录").JSON(c)
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}
