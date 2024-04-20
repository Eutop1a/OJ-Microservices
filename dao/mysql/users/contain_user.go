package users

import (
	"OnlineJudge/dao/mysql"
)

func ContainUser(username string) bool {

	var user mysql.User

	mysql.DB.First(&user, "username = ?", username)

	if user.UserName != username {
		//fmt.Println(user.Username, "!=", username)
		return false
	}
	//fmt.Println(user.Username, "!=", username)
	return true
}
