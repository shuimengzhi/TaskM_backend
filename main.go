package main

import (
	"flag"
	"os"
	"taskm/core"
	//_ "taskm/docs"
	"taskm/router"
)

func main() {
	envPath, _ := os.Getwd()
	envPath = envPath + "/.env"

	//如果有输入.env文件位置，那么优先使用输入参数信息
	inputEnvPath := ""
	flag.StringVar(&inputEnvPath, "e", "", ".env文件位置")
	flag.Parse()
	if inputEnvPath != "" {
		envPath = inputEnvPath
	}
	core.LoadCore(envPath)
	r := router.NewRouter()
	r.Run(":4000")
}
