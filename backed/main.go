package main

import (
	"context"
	"fmt"
	"myApp/global"
	"myApp/initialize"
	"myApp/pkg/errors"
	"myApp/routes"
	"myApp/utils"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	UseValidator "myApp/validators"

	_ "myApp/swagger/docs"

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
	err := initialize.InitDB()
	if err != nil {
		zap.S().Panic(err, errors.ParseCoder(err).Code())
	}
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
	freeport, err := utils.GetFreePort()

	if err != nil {
		zap.S().Panic(err.Error())
	}
	////启动服务器
	srv := &http.Server{Addr: fmt.Sprintf(":%d", freeport), Handler: r}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Errorw("listenAndServe err", "error", err)
		}
	}()
	//优雅关停服务
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(shutdownCtx)
}
