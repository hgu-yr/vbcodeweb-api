package handler

import (
	"github.com/gin-gonic/gin"

	"vbCodeWeb/common"
	"vbCodeWeb/request"
	"vbCodeWeb/service"
)

// AuthController 鉴权接口。
type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Login 管理员登录
func (ac *AuthController) Login(c *gin.Context) {
	var req request.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.Err(c, err.Error())
		return
	}
	result, err := ac.authService.Login(&req)
	if err != nil {
		common.Err(c, err.Error())
		return
	}
	common.OK(c, result)
}

// Profile 返回当前登录用户信息
func (ac *AuthController) Profile(c *gin.Context) {
	common.OK(c, gin.H{
		"user_id":  c.GetUint("userID"),
		"username": c.GetString("username"),
		"role":     c.GetString("role"),
	})
}
