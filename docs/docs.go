// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Eutop1a",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/sort": {
            "post": {
                "description": "获取所有用户并按分数降序排序",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "根据用户分数降序排序",
                "parameters": [
                    {
                        "type": "string",
                        "description": "执行操作的用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取并排序用户列表",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseSort"
                        }
                    },
                    "403": {
                        "description": "Token 已超时",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "数据库查询错误",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseError"
                        }
                    }
                }
            }
        },
        "/changeScore": {
            "post": {
                "description": "通过用户名和新分数的表单数据，增加用户的总分数，并返回新的总分数",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "增加用户分数",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "要增加的分数",
                        "name": "newscore",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "增加分数成功，返回消息和新的总分数",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseAddScore"
                        }
                    },
                    "403": {
                        "description": "Token 已超时",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "数据库查询或保存出错",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseError"
                        }
                    }
                }
            }
        },
        "/judge": {
            "post": {
                "description": "通过用户身份、问题ID和代码内容，进行代码评测，并返回评测结果",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "代码评测",
                "parameters": [
                    {
                        "type": "string",
                        "description": "执行操作的用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "要评测的问题ID",
                        "name": "problem",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "要评测的代码内容",
                        "name": "code",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功进行代码评测",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseMsg"
                        }
                    },
                    "403": {
                        "description": "Token 已超时",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseMsg"
                        }
                    }
                }
            }
        },
        "/list": {
            "post": {
                "description": "获取所有问题的列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取问题列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "要获取问题列表的用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取问题列表",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseProblems"
                        }
                    },
                    "403": {
                        "description": "Token 已超时",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseMsg"
                        }
                    },
                    "500": {
                        "description": "获取问题列表出错",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseMsg"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "用户登录接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "user_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "验证码",
                        "name": "code",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功",
                        "schema": {
                            "$ref": "#/definitions/controller._LoginSuccess"
                        }
                    },
                    "400": {
                        "description": "用户名不存在或验证码错误",
                        "schema": {
                            "$ref": "#/definitions/controller._LoginError"
                        }
                    },
                    "401": {
                        "description": "验证码过期",
                        "schema": {
                            "$ref": "#/definitions/controller._LoginError"
                        }
                    },
                    "403": {
                        "description": "密码错误",
                        "schema": {
                            "$ref": "#/definitions/controller._LoginError"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/controller._LoginError"
                        }
                    }
                }
            }
        },
        "/problem/:id": {
            "post": {
                "description": "通过问题ID和管理员身份，获取指定问题的详细信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取问题详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "执行操作的管理员用户名",
                        "name": "_",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "要获取的问题ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取问题详情",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseQuestionDetail"
                        }
                    },
                    "403": {
                        "description": "获取问题详情出错",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseMsg"
                        }
                    }
                }
            }
        },
        "/problem/add": {
            "post": {
                "description": "通过管理员身份，使用表单数据添加新的问题，并返回操作结果",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "添加问题",
                "parameters": [
                    {
                        "type": "string",
                        "description": "执行操作的管理员用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "问题ID",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "问题标题",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "问题描述",
                        "name": "lore",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "问题标准输入",
                        "name": "input",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "问题标准输出",
                        "name": "output",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "问题提示",
                        "name": "tips",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功添加问题",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseAddProblems"
                        }
                    },
                    "403": {
                        "description": "解析表单数据出错",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseMsg"
                        }
                    }
                }
            }
        },
        "/problem/file/:id": {
            "post": {
                "description": "通过问题ID，获取指定问题的文件列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取问题文件列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "要获取文件列表的问题ID",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功获取问题文件列表",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseGetData"
                        }
                    },
                    "403": {
                        "description": "读取文件列表出错",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseDataString"
                        }
                    }
                }
            }
        },
        "/problem/file/add/:id": {
            "post": {
                "description": "通过管理员身份，使用表单数据上传问题的输入和输出文件，并返回操作结果",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "上传问题文件",
                "parameters": [
                    {
                        "type": "string",
                        "description": "执行操作的管理员用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "要上传文件的问题ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "问题输入文件（.in）",
                        "name": "input",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "问题输出文件（.out）",
                        "name": "output",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功上传问题文件",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseMsg"
                        }
                    },
                    "400": {
                        "description": "接收或保存文件出错",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseMsg"
                        }
                    },
                    "403": {
                        "description": "Token 已超时或用户非管理员",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponseMsg"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "用户注册接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "user_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "验证码",
                        "name": "code",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "注册成功",
                        "schema": {
                            "$ref": "#/definitions/controller._RegisterSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller._RegisterError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/controller._RegisterError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller._RegisterError"
                        }
                    }
                }
            }
        },
        "/send-email-code": {
            "post": {
                "description": "发送验证码接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "发送验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "发送验证码成功",
                        "schema": {
                            "$ref": "#/definitions/controller._SendCodeSuccess"
                        }
                    },
                    "400": {
                        "description": "邮箱格式错误",
                        "schema": {
                            "$ref": "#/definitions/controller._SendCodeError"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/controller._SendCodeError"
                        }
                    }
                }
            }
        },
        "/users/{user_id}": {
            "get": {
                "description": "获取用户详细信息接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取用户详细信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "$ref": "#/definitions/controller._GetUserDetailSuccess"
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/controller._GetUserDetailError"
                        }
                    },
                    "403": {
                        "description": "没有此用户ID",
                        "schema": {
                            "$ref": "#/definitions/controller._GetUserDetailError"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/controller._GetUserDetailError"
                        }
                    }
                }
            },
            "put": {
                "description": "更新用户详细信息接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "更新用户详细信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户密码",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "用户邮箱",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "$ref": "#/definitions/controller._UpdateUserDetailSuccess"
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/controller._UpdateUserDetailError"
                        }
                    },
                    "403": {
                        "description": "没有此用户ID or 验证码错误",
                        "schema": {
                            "$ref": "#/definitions/controller._UpdateUserDetailError"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/controller._UpdateUserDetailError"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除用户接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "获取成功",
                        "schema": {
                            "$ref": "#/definitions/controller._DeleteUserSuccess"
                        }
                    },
                    "400": {
                        "description": "参数错误",
                        "schema": {
                            "$ref": "#/definitions/controller._DeleteUserError"
                        }
                    },
                    "403": {
                        "description": "没有此用户ID",
                        "schema": {
                            "$ref": "#/definitions/controller._DeleteUserError"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/controller._DeleteUserError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller._DeleteUserError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "format": "string",
                    "example": "delete user error"
                }
            }
        },
        "controller._DeleteUserSuccess": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "format": "string",
                    "example": "success delete user"
                }
            }
        },
        "controller._GetUserDetailError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "format": "string",
                    "example": "get user detail error"
                }
            }
        },
        "controller._GetUserDetailSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/mysql.User"
                },
                "msg": {
                    "type": "string",
                    "format": "string",
                    "example": "success get user detail"
                }
            }
        },
        "controller._LoginError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "format": "string",
                    "example": "invalidate email format"
                }
            }
        },
        "controller._LoginSuccess": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "format": "string",
                    "example": "login successfully"
                }
            }
        },
        "controller._RegisterError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "format": "string",
                    "example": "register error"
                }
            }
        },
        "controller._RegisterSuccess": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "format": "string",
                    "example": "registration successful"
                },
                "token": {
                    "type": "string",
                    "format": "token string",
                    "example": "bearer token"
                }
            }
        },
        "controller._ResponseAddProblems": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "controller._ResponseAddScore": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                },
                "new_score": {
                    "type": "integer"
                }
            }
        },
        "controller._ResponseDataString": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                }
            }
        },
        "controller._ResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "controller._ResponseGetData": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "controller._ResponseMsg": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        },
        "controller._ResponseProblems": {
            "type": "object"
        },
        "controller._ResponseQuestionDetail": {
            "type": "object"
        },
        "controller._ResponseSort": {
            "type": "object"
        },
        "controller._SendCodeError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "format": "string",
                    "example": "invalidate email format"
                }
            }
        },
        "controller._SendCodeSuccess": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "format": "string",
                    "example": "send verification code successfully"
                }
            }
        },
        "controller._UpdateUserDetailError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "format": "string",
                    "example": "update user information error"
                }
            }
        },
        "controller._UpdateUserDetailSuccess": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "format": "string",
                    "example": "success update user information"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "mysql.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "last_login_data": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "registration_date": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "user_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:65533",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "OnlineJudge",
	Description:      "Refactoring",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
