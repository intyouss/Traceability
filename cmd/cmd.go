package cmd

import (
	"fmt"
	"github.com/intyouss/Traceability/config"
	"github.com/intyouss/Traceability/router"
)

func Start() {
	config.InitConfig()
	router.InitRouter()
}

func Clean() {
	fmt.Println("clean")
}
