package users

import (
	"OnlineJudge/dao/mysql"
)

func JudgePWD(username string, pwd string) bool {

	var user mysql.User
	mysql.DB.First(&user, "username = ?", username)

	if user.UserName != username {
		return false
	}

	return user.Password == pwd

}
