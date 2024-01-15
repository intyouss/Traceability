package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/intyouss/Traceability/docs"
	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/middleware"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type IFnRegisterRoute = func(defaultGroup *gin.RouterGroup, authGroup *gin.RouterGroup)

// 注册路由列表
var (
	routes []IFnRegisterRoute
)

// RegisterRoute 添加路由至列表
func RegisterRoute(fn IFnRegisterRoute) {
	if fn == nil {
		return
	}
	routes = append(routes, fn)
}

// InitRouter 初始化路由
func InitRouter() {
	// 监听 ctrl + c和应用退出信号
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	r := gin.Default()
	r.Use(middleware.Cors())
	defaultGroup := r.Group("/api/v1/public")
	authGroup := r.Group("api/v1/")
	authGroup.Use(middleware.Auth())

	InitBaseRoutes()

	//// 注册自定义校验器
	//registerValidation()

	// 初始化所有注册路由
	for _, fnRegisterRoute := range routes {
		fnRegisterRoute(defaultGroup, authGroup)
	}

	// 集成Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := viper.GetString("server.port")
	if port == "" {
		port = "8090" // default
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	go func() {
		ln, err := net.Listen("tcp", server.Addr)
		if err != nil {
			global.Logger.Error(fmt.Sprintf("Start Server Error:%s", err.Error()))
			return
		}
		global.Logger.Infof("Listening and serving HTTP on %s", server.Addr)
		if err := server.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.Logger.Error(fmt.Sprintf("Start Server Error:%s", err.Error()))
			return
		}
	}()

	<-ctx.Done()

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("Stop Server Error: %s", err.Error()))
		return
	}

	global.Logger.Info("Stop Server Success")
}

func InitBaseRoutes() {
	InitUserRoutes()
}

//func registerValidation() {
//	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
//		_ = v.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
//			if mobile, ok := fl.Field().Interface().(string); ok {
//				if mobile != "" && strings.Index(mobile, "1") == 0 {
//					return true
//				}
//			}
//			return false
//		})
//	}
//}
