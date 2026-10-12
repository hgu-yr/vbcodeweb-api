package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// parseID 解析路径参数 id，失败返回 false。
func parseID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0, false
	}
	return uint(id), true
}
