package v1

import (
	"linglong/models"
	"linglong/pkg/e"
	"linglong/pkg/util"
	"net/http"
	"github.com/gin-gonic/gin"
)

// LoginRequest 登录请求参数
type LoginRequest struct {
	Username string `form:"username" binding:"required,min=3,max=10"`
	Password string `form:"password" binding:"required,min=3,max=16"`
}

// Login 处理用户登录
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  "参数错误",
			"data": nil,
		})
		return
	}

	// 验证用户名密码（核心修复：调用模型层校验）
	if models.CheckAuth(req.Username, req.Password) {
		// 生成 JWT Token
		token, err := util.GenerateToken(req.Username)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": e.ERROR,
				"msg":  "Token 生成失败",
				"data": nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  "登录成功",
			"data": gin.H{"token": token},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR_AUTH,
			"msg":  "用户名或密码错误",
			"data": nil,
		})
	}
}
