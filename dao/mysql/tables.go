package mysql

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// User 用户登录基本信息
type User struct {
	Model
	UserID           int64     `gorm:"type:bigint;primaryKey;column:userID" json:"user_id"`
	UserName         string    `gorm:"type:varchar(255);not null;column:userName" json:"user_name"`
	Password         string    `gorm:"type:varchar(255);not null;column:password" json:"password"`
	Email            string    `gorm:"type:varchar(255);unique;not null;column:email" json:"email"`
	Score            int       `gorm:"type:int;not null;column:score" json:"score"`
	RegistrationDate time.Time `gorm:"type:timestamp;not null;column:registrationDate" json:"registration_date"`
	LastLoginData    time.Time `gorm:"type:timestamp;column:lastLoginData" json:"last_login_data"`
}

type UserSolved struct {
	gorm.Model
	Username string
	Solved   string
}

type ProblemList struct {
	gorm.Model
	ProblemID          int    //问题ID
	ProblemScore       int    //问题分值
	MaxRuntime         int    `json:"maxRuntime" gorm:"column:maxRuntime;type:int(11);"` //最大运行时长,单位为毫秒
	MaxMemory          int    `json:"maxMemory" gorm:"column:maxMemory;type:int(11)"`
	ProblemTitle       string //问题标题
	ProblemLore        string //问题内容
	ProblemStandardIn  string //问题输入样例
	ProblemStandardOut string //问题输出样例
	ProblemTips        string //问题提示
	Author             string //发布人

}

type Problems struct {
	gorm.Model
	ProblemID    int    //问题ID
	ProblemTitle string //问题标题
	Author       string //发布人
}

type Admins struct {
	gorm.Model
	Username string
}

// TableName gorm自动改表名
func (u *User) TableName() string {
	return "User"
}

func (p *Problems) TableName() string {
	return "Problems"
}

func (s *UserSolved) TableName() string {
	return "UserSolved"
}

func (j *ProblemList) TableName() string {
	return "ProblemList"
}

func (p *Admins) TableName() string {
	return "Admins"
}
