package users

import (
	"OnlineJudge/dao/mysql"
	"crypto/md5"
	"fmt"
)

const Title = "Eutop1a"

func NewUser(username, password string) error {

	pwd := fmt.Sprintf("%x", md5.Sum([]byte(password+Title)))

	mysql.DB.Create(&mysql.User{
		UserName: username,
		Password: pwd,
	})

	return nil
}
