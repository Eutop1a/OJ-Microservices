package main

import (
	"OnlineJudge/dao/mysql"
	"OnlineJudge/dao/redis"
	"OnlineJudge/logger"
	"OnlineJudge/routes"
	"OnlineJudge/setting"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title OnlineJudge
// @version 1.0
// @description Refactoring
// @termsOfService http://swagger.io/terms/

// @contact.name Eutop1a
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:65533
// @BasePath /api/v1
func main() {
	// 1、加载配置文件
	if err := setting.Init(); err != nil {
		fmt.Printf("init setting failed, err: %v\n", err)
		return
	}
	// 2、初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}
	defer zap.L().Sync()

	// 3、初始化MySQL连接
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err: %v\n", err)
		return
	}
	//defer mysql.Close()

	// 4. init Redis connection
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("Init redis failed, err: %v\n", err)
		return
	}

	// 5、注册路由
	r := routes.SetUp(setting.Conf.Mode)

	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("Run server failed, %v\n", err)
		return
	}

	// 6、启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown:", zap.Error(err))
	}
	zap.L().Info("Server exiting")

}
