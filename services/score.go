package services

import (
	"OnlineJudge/dao/mysql"
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

type ChangeScore struct {
	UserName  string `form:"username"`
	ProblemId string `form:"problem"`
}

func (score *ChangeScore) AddScore() (int, string, int) {
	// 从请求中获取用户名和新分数
	username := score.UserName
	problemId := score.ProblemId

	// 通过用户名查询用户信息
	var user mysql.User
	err := mysql.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		// 数据库查询出错
		zap.L().Error(fmt.Sprintf("database select %s error", username), zap.Error(err))
		return http.StatusInternalServerError, err.Error(), 0
	}

	// 获取题目分数
	var problem mysql.ProblemList
	err = mysql.DB.Where("problem_id = ?", problemId).First(&problem).Error
	if err != nil {
		// 数据库查询出错
		zap.L().Error(fmt.Sprintf("database get problem id: %s error", problemId), zap.Error(err))
		return http.StatusInternalServerError, err.Error(), 0
	}

	// 解析新分数并更新用户总分数

	user.Score += problem.ProblemScore

	// 保存更新后的用户信息
	err = mysql.DB.Save(&user).Error
	if err != nil {
		// 数据库保存出错
		zap.L().Error(fmt.Sprintf("database save %s error", username), zap.Error(err))

		return http.StatusInternalServerError, err.Error(), 0
	}

	// 返回成功响应
	return http.StatusOK, "", user.Score

}

func (score *ChangeScore) SortByScore() (int, string, []mysql.User) {
	// 从请求中获取执行操作的用户名
	//username := score.UserName

	// 查询并按分数降序排序用户列表
	var users []mysql.User
	err := mysql.DB.Order("score desc").Find(&users).Error
	if err != nil {
		zap.L().Error("database order error", zap.Error(err))
		return http.StatusInternalServerError, err.Error(), nil
	}

	return http.StatusOK, "", users
}
