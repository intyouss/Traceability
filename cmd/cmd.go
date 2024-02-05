package cmd

import (
	"github.com/intyouss/Traceability/config"
	"github.com/intyouss/Traceability/global"
	"github.com/intyouss/Traceability/router"
)

func Start() {
	var err error
	// 初始化系统配置
	if err = config.InitConfig(); err != nil {
		panic(err)
	}
	// 初始化日志
	global.Logger = config.InitLogger()
	// 初始化数据库
	if global.DB, err = config.InitDB(); err != nil {
		panic(err)
	}
	// 初始化Redis
	if global.Redis, err = config.InitRedis(); err != nil {
		panic(err)
	}
	// 初始化OSS
	if global.OSS, err = config.InitOSS(); err != nil {
		panic(err)
	}
	// 初始化路由
	router.InitRouter()
}

func Clean() {
}
