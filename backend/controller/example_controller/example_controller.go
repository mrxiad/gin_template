package example_controller

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	// "template/constants"
	// daoImpl "template/dao/impl"
	// "template/db/mysql"
	// "template/model"
	// "template/model/params"
	// serviceImpl "template/service/example_service"
)

type ExampleController struct {
}


// 注册请求参数结构体
type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 注册响应数据结构体
type SignUpResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		UserID string `json:"userId"`
	} `json:"data"`
}

// SignUpHandler 处理注册请求
func (o ExampleController) SignUpHandler(c *gin.Context) {
	// 1. 解析请求参数
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{"code": 400, "msg": err.Error(), "data": gin.H{}})
		return
	}

	// 2. 校验参数
	if req.Username == "" || req.Password == "" {
		c.JSON(200, gin.H{"code": 400, "msg": "用户名和密码不能为空", "data": gin.H{}})
		return
	}

	// 3. 处理业务逻辑，模拟注册成功
	resp := SignUpResponse{
		Code: 1001,
		Msg:  "注册成功",
	}
	resp.Data.UserID = "123456"

	// 4. 返回结果
	c.JSON(200, gin.H{"code": resp.Code, "msg": resp.Msg, "data": resp.Data})
}


// 登录请求参数结构体
type LoginRequest struct {
	UserID string `json:"userId"`
	Password string `json:"password"`
}

// 登录响应数据结构体
type LoginResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct{} `json:"data"`
}

// loginHandler 处理登录请求
func (o ExampleController) LoginHandler(c *gin.Context) {
	// 1. 解析请求参数
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, LoginResponse{Code: 1004, Msg: "登录失败", Data: struct{}{}})
		return
	}

	// 2. 校验参数
	if req.UserID == "" || req.Password == "" {
		c.JSON(200, LoginResponse{Code: 1004, Msg: "登录失败", Data: struct{}{}})
		return
	}

	// 3. 处理业务逻辑，模拟登录成功
	resp := LoginResponse{
		Code: 1003,
		Msg:  "登录成功",
		Data: struct{}{},
	}

	// 4. 返回结果
	c.JSON(200, resp)
}