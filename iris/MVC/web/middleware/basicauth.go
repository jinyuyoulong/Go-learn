package middleware

import "github.com/kataras/iris/middleware/basicauth"

// 简单的授权验证
var BasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		// 需要校验的 用户名和密码
		"admin": "password",
	},
})
