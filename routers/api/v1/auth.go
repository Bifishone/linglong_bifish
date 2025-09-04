package v1

import (
	"linglong/models"
	"linglong/pkg/e"
	"linglong/pkg/util"
	"net/http"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// 登录请求参数
type LoginRequest struct {
	Username string `form:"username" valid:"Required;MinSize(3);MaxSize(10)"`
	Password string `form:"password" valid:"Required;MinSize(3);MaxSize(16)"`
}

// @Summary 用户登录
// @Produce  json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} util.Response
// @Failure 400 {object} util.Response
// @Failure 500 {object} util.Response
// @Router /auth [post]
func Login(c *gin.Context) {
	// 获取请求参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 参数验证
	valid := validation.Validation{}
	valid.Required(username, "username").Message("用户名不能为空")
	valid.Required(password, "password").Message("密码不能为空")
	valid.MinSize(username, 3, "username").Message("用户名长度不能小于3")
	valid.MaxSize(username, 10, "username").Message("用户名长度不能大于10")
	valid.MinSize(password, 3, "password").Message("密码长度不能小于3")
	valid.MaxSize(password, 16, "password").Message("密码长度不能大于16")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		// 验证用户名密码是否正确
		if models.CheckAuth(username, password) {
			// 生成JWT令牌
			token, err := util.GenerateToken(username)
			if err != nil {
				code = e.ERROR
			} else {
				code = e.SUCCESS
				c.JSON(http.StatusOK, gin.H{
					"code": code,
					"msg":  e.GetMsg(code),
					"data": gin.H{"token": token},
				})
				return
			}
		} else {
			code = e.INVALID_PASS // 对应402错误码（无效密码）
		}
	}

	// 返回错误信息
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
