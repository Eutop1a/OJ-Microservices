package middlewares

import (
	"OnlineJudge/pkg"
	"github.com/gin-gonic/gin"
)

// AuthorizationMiddleWare 认证中间件

func AuthorizationMiddleWare(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	var username string
	var succ bool
	if token[:6] == "Bearer" {
		token = token[7:]
	}
	if username, succ = pkg.JudgeAccessToken(token); succ {

		c.Request.ParseForm()
		c.Request.Form.Set("username", username)

		//// 获取修改后的username值
		//newUsername := c.Request.Form.Get("username")
		//fmt.Println("New username:", newUsername)
		//c.Params["username"] = username
		c.Next()
	} else {
		//delete(c.Params, "username")
		c.Request.ParseForm()
		c.Request.Form.Del("username")
		c.Next()
	}
}
