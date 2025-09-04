package routers

import (
	"linglong/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	apiv1 := r.Group("/api/v1")
	{
		// 登录接口（核心修复：注册登录路由）
		apiv1.POST("/auth", v1.Login)
	}
}
