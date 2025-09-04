package v1

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(api *gin.RouterGroup) {
	// 登录接口
	api.POST("/auth", Login)
	// 其他接口...
	api.POST("/setting/editpass", EditPass)
	api.POST("/finger/test", TestFinger)
}
