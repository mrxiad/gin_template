package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"template/config"
	// "template/db/mysql"
	"template/logger"
	"template/router"
	"time"
)

func main() {
	var err error

	//初始化配置
	err = config.Init("./config/config.yaml")
	if err != nil {
		fmt.Println("初始化配置失败", err)
		return
	}

	//初始化日志
	err = logger.Init(config.Conf.LogConfig, config.Conf.Mode)
	if err != nil {
		fmt.Println("初始化日志失败", err)
		return
	}

	//初始化mysql
	// err = mysql.Init(config.Conf.MySQLConfig)
	// if err != nil {
	// 	fmt.Println("初始化mysql失败", err)
	// 	return
	// }
	// defer mysql.Close()

	//初始化redis
	//err = redis.Init(config.Conf.RedisConfig)
	//if err != nil {
	//	fmt.Println("初始化redis失败", err)
	//	return
	//}
	//defer redis.Close()

	//加载路由
	r := router.SetupRouter(config.Conf.Mode)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Conf.Port),
		Handler: r, //因为gin.Engine实现了ServeHTTP接口，所以可以直接赋值
	}

	//开启服务
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	//优雅关机
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	// 接收 syscall.SIGINT 和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("正在关闭服务器...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("服务器关闭:", err)
	}
	fmt.Println("服务器退出")
}
