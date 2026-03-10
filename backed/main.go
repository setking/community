package main

import (
	"fmt"

	"myApp/global"
	"myApp/initialize"
	"myApp/routes"

	UseValidator "myApp/validators"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func main() {
	// 加载配置
	initialize.InitConfig()
	////初始化日志
	initialize.InitLogger()
	////初始化mysql连接
	initialize.InitDB()
	//初始化redis连接
	initialize.InitRedis()
	defer initialize.CloseRedis()
	//初始化snowflake
	initialize.InitSnowflake(1)
	////注册路由
	r := routes.Routers(global.ServerConfig.MyAppInfo.Mode)
	// 初始化翻译
	if err := initialize.InitTrans(global.ServerConfig.MyAppInfo.Local); err != nil {
		zap.S().Panic(err)
	}
	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("custom_email", UseValidator.ValidateEmail)
		_ = v.RegisterValidation("custom_direction", UseValidator.ValidateDirection)
		_ = v.RegisterTranslation("custom_email", global.Trans, func(ut ut.Translator) error {
			return ut.Add("custom_email", "{0} 邮箱格式不正确", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("custom_email", fe.Field())
			return t
		})
		_ = v.RegisterTranslation("custom_direction", global.Trans, func(ut ut.Translator) error {
			return ut.Add("custom_direction", "{0} 值不正确", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("custom_direction", fe.Field())
			return t
		})
	}
	////启动服务器
	err := r.Run(fmt.Sprintf(":%d", global.ServerConfig.MyAppInfo.Port))
	if err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

	//
	////// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	//quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	////// kill 默认会发送 syscall.SIGTERM 信号
	////// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	////// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	////// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	//<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	//zap.L().Info("Shutdown Server ...")
	////// 创建一个5秒超时的context
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	////// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	//if err := srv.Shutdown(ctx); err != nil {
	//	zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	//}
	//
	//zap.L().Info("Server exiting")
}
