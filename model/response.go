package model

import (
	"github.com/gin-gonic/gin"
	"template/constants"
)

type ResponseCode struct {
	Code constants.ResponseCodeType `json:"code"`
	Desc string                     `json:"desc"`
}

type ResponseStruct struct {
	ResponseCode
	Data interface{} `json:"data"`
}

// ResponseWithData 带data的返回结果
func ResponseWithData(c *gin.Context, response ResponseStruct) {
	////添加描述
	//response.ResponseCode.Desc = response.ResponseCode.Code.String()
	//返回结果
	c.JSON(200, response)
}

// ResponseWithoutData 不带data的返回结果
func ResponseWithoutData(c *gin.Context, code constants.ResponseCodeType) {
	//构造返回结果
	response := ResponseStruct{
		ResponseCode: ResponseCode{
			Code: code,
			Desc: code.String(),
		},
		Data: nil,
	}
	//返回结果
	ResponseWithData(c, response)
}
