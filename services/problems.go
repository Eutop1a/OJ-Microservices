package services

import (
	"OnlineJudge/dao/mysql"
	"OnlineJudge/dao/mysql/users"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Problem struct {
	UserName string `form:"username" json:"username"`
	Id       string `form:"id" json:"id"`
	Title    string `form:"title" json:"title"`
	Lore     string `form:"lore" json:"lore"`
	Input    string `form:"input" json:"input"`
	Output   string `form:"output" json:"output"`
	Tips     string `form:"tips" json:"tips"`
	Score    int    `form:"score" json:"score"`
}

func (problem *Problem) GetList() (int, string, []mysql.Problems) {

	// 查询问题列表
	var datas []mysql.Problems
	result := mysql.DB.Find(&datas)

	if result.Error != nil {
		// 获取问题列表出错
		zap.L().Error("get list error", zap.Error(result.Error))

		return http.StatusInternalServerError, "get list error", nil

	} else {
		// 成功获取问题列表
		return http.StatusOK, "", datas
	}
}

func (problem *Problem) AddProblem() (int, string) {

	// 判断用户是否为管理员
	if !users.JudgeAdmin(problem.UserName) {

		return http.StatusForbidden, "Forbidden"
	}

	// 将问题添加到数据库
	ID, _ := strconv.Atoi(problem.Id)
	mysql.DB.Create(&mysql.ProblemList{
		ProblemID:          ID,
		ProblemTitle:       problem.Title,
		ProblemLore:        problem.Lore,
		ProblemStandardIn:  problem.Input,
		ProblemStandardOut: problem.Output,
		ProblemTips:        problem.Tips,
		Author:             problem.UserName,
		ProblemScore:       problem.Score,
	})

	mysql.DB.Create(&mysql.Problems{
		ProblemTitle: problem.Title,
		ProblemID:    ID,
		Author:       problem.UserName,
	})

	return http.StatusOK, "Ok"
}

func (problem *Problem) GetProblem(id string) (int, string, mysql.ProblemList) {

	// 查询指定问题的详细信息
	var data mysql.ProblemList
	mysql.DB.First(&data, "problem_id = ?", id)
	if data.ProblemTitle == "" {
		// 获取问题详情出错
		zap.L().Error("get problem error")

		return http.StatusBadRequest, "get problem error", mysql.ProblemList{}
	}
	return http.StatusOK, "", data
}
