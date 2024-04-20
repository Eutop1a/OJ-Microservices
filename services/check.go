package services

import (
	"OnlineJudge/dao/mysql/users"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type Check struct {
	UserName string `form:"username"`
	NewScore int    `form:"newscore"`
}

func (check *Check) CheckUserName() (int, string) {
	// 判断用户名是否已存在
	status := users.ContainUser(check.UserName)
	if status {
		// 用户名已存在
		zap.L().Error(fmt.Sprintf("username %s already exists", check.UserName))
		return http.StatusForbidden, "UserExist"
	} else {
		// 用户名不存在
		return http.StatusOK, "notExist"
	}
}
