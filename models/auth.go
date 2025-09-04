package models

// 模拟用户数据（实际可对接数据库，此处为测试用默认账号）
var mockUsers = map[string]string{
	"linglong": "123456", // 默认用户名：linglong，密码：123456
}

// CheckAuth 验证用户名密码
func CheckAuth(username, password string) bool {
	pwd, exists := mockUsers[username]
	return exists && pwd == password
}
