package router

import (
	"github.com/gin-gonic/gin"
	"template/router/example_router"
)

// SetupRouter 路由
// 这里初始化所有路由，返回一个gin引擎
func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()

	//配置日志
	//r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//初始化example路由
	example_router.InitExampleRouter(r)

	
	//配置各个文件夹的路由
	return r
}
