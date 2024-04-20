package users

import (
	"OnlineJudge/dao/mysql"
)

func JudgeAdmin(username string) bool {

	var admin mysql.Admins

	mysql.DB.First(&admin, "username = ?", username)

	return admin.Username == username

}

func RegAdmin(username string) int {
	var existingAdmin mysql.Admins

	// 检查用户名是否已经存在
	err := mysql.DB.First(&existingAdmin, "username = ?", username).Error
	if err == nil {
		// 用户名已经存在，不进行创建
		return 1
	}

	// 创建新的管理员记录
	newAdmin := mysql.Admins{
		Username: username,
	}

	err = nil
	err = mysql.DB.Create(&newAdmin).Error
	if err != nil {
		// 内部错误
		return 2
	}
	// 返回true表示成功创建管理员
	return 0
}
