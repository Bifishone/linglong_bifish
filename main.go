package main

import (
	"linglong/global"
	"linglong/routers"
	"linglong/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 注册跨域中间件（核心修复：允许前端跨域请求）
	r.Use(middleware.Cors())

	// 初始化路由
	routers.InitRouter(r)

	// 启动服务（端口与前端请求匹配）
	r.Run(":8080")
}
