package example_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"template/constants"
	daoImpl "template/dao/impl"
	"template/db/mysql"
	"template/model"
	"template/model/params"
	serviceImpl "template/service/example_service"
)

type ExampleController struct {
}

// SingUpHandler 必须返回json，只可以返回一次
func (o ExampleController) SingUpHandler(c *gin.Context) {
	//1.校验参数
	exampleRequestParam := params.ExampleRequestParam{}
	if err := c.ShouldBindJSON(&exampleRequestParam); err != nil {
		//返回错误
		fmt.Println(err)
		//发生错误返回，c+错误码
		model.ResponseWithoutData(c, constants.ResponseCodeParamError)
		return
	}

	//2.调用service，创建一个service实例
	exampleService := serviceImpl.ExampleService{
		//依赖注入
		ExampleDao: daoImpl.NewExampleDaoImpl(mysql.Db),
	}

	//注意：这个结果是在db/table下的表struct
	exampleTable, err := exampleService.QueryExampleById(uint(exampleRequestParam.Id))
	if err != nil {
		//返回错误,可能数据库里没东西
		fmt.Println(err)
		model.ResponseWithoutData(c, constants.ResponseCodeServerError)
		return
	}
	//3.返回结果
	fmt.Println(exampleTable)
	model.ResponseWithData(c, model.ResponseStruct{
		ResponseCode: model.ResponseCode{
			Code: constants.ResponseCodeSuccess,
			Desc: constants.ResponseCodeSuccess.String(),
		},
		Data: exampleTable,
	})
	return
}
