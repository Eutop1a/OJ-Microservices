package controller

import "OnlineJudge/dao/mysql"

type _ResponseCode struct {
	Code int `json:"code"`
}
type _ResponseError struct {
	Error string `json:"error"`
}

type _ResponseMsg struct {
	Msg string `json:"msg"`
}

type _ResponseToken struct {
	Token string `json:"Token"`
}

type _ResponseReg struct {
	_ResponseToken
	EndTime string `json:"end_time"`
}

type _ResponseRegErr struct {
	_ResponseMsg
	Err string `json:"err"`
}

type _ResponseSort struct {
	//Data []User `json:"sort"`
}

type _ResponseAddScore struct {
	_ResponseMsg
	NewScore int `json:"new_score"`
}

type _ResponseUsername struct {
	Username string `json:"username"`
}

type _ResponseProblems struct {
	//Msg []Problems `json:"judge"`
}

type _ResponseAddProblems struct {
	Code string `json:"code"`
}

type _ResponseDataString struct {
	Data string `json:"data"`
}
type _ResponseGetData struct {
	Data []string `json:"data"`
}

type _ResponseQuestionDetail struct {
	//Data ProblemList `json:"data"`
}

// _RegisterSuccess 注册成功返回结构体
type _RegisterSuccess struct {
	Token string `json:"token" example:"bearer token" format:"token string"`
	Msg   string `json:"msg" example:"registration successful" format:"string"`
}

// _RegisterError 针对注册失败情况的返回结构体
type _RegisterError struct {
	Error string `json:"error" example:"register error" format:"string"`
}

// _LoginSuccess 登录成功返回结构体
type _LoginSuccess struct {
	Msg string `json:"msg" example:"login successfully" format:"string"`
}

// _LoginError 登录失败返回结构体
type _LoginError struct {
	Error string `json:"error" example:"invalidate email format" format:"string"`
}

// _SendCodeSuccess 发送验证码成功返回结构体
type _SendCodeSuccess struct {
	Msg string `json:"msg" example:"send verification code successfully" format:"string"`
}

// _SendCodeError 发送验证码失败返回结构体
type _SendCodeError struct {
	Error string `json:"error" example:"invalidate email format" format:"string"`
}

// _GetUserDetailError 获取用户信息成功返回结构体
type _GetUserDetailSuccess struct {
	Msg  string     `json:"msg" example:"success get user detail" format:"string"`
	Data mysql.User `json:"data"`
}

// _GetUserDetailError 获取用户信息失败返回结构体
type _GetUserDetailError struct {
	Error string `json:"error" example:"get user detail error" format:"string"`
}

// _DeleteUserSuccess 获取用户信息成功返回结构体
type _DeleteUserSuccess struct {
	Msg string `json:"msg" example:"success delete user" format:"string"`
}

// _DeleteUserError 获取用户信息失败返回结构体
type _DeleteUserError struct {
	Error string `json:"error" example:"delete user error" format:"string"`
}

// _DeleteUserSuccess 获取用户信息成功返回结构体
type _UpdateUserDetailSuccess struct {
	Msg string `json:"msg" example:"success update user information" format:"string"`
}

// _DeleteUserError 获取用户信息失败返回结构体
type _UpdateUserDetailError struct {
	Error string `json:"error" example:"update user information error" format:"string"`
}
