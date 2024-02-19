package example_router

import (
	"github.com/gin-gonic/gin"
	example_controller "template/controller/example_controller"
)

func InitExampleRouter(r *gin.Engine) {

	//开放了/example_service/id这个api，然后调用controller层
	exampleGroup := r.Group("/example")
	{
		exampleGroup.POST("/id", example_controller.ExampleController{}.SingUpHandler)
	}
}
