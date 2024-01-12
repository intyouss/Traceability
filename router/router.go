package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/intyouss/Traceability/docs"
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
	defaultGroup := r.Group("/api/v1/public")
	authGroup := r.Group("api/v1/")

	InitBaseRoutes()
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
			panic(err.Error())
			return
		}
		fmt.Printf("Listening and serving HTTP on %s\n", server.Addr)
		if err := server.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println(fmt.Sprintf("Start Server Error:%s", err.Error()))
			return
		}
	}()

	<-ctx.Done()

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Stop Server Error: %s", err.Error())
		return
	}

	fmt.Println("Stop Server Success")
}

func InitBaseRoutes() {
	InitUserRoutes()
}
