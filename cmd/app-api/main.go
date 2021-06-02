package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"tinkgo/internal/app-api/pkg/core"
	"tinkgo/internal/app-api/router"
	"tinkgo/pkg/tinkgo/api"
)

var env = kingpin.Flag("env", "Set run environment, options: prod / test / dev").Default("dev").String()

func main()  {
	// 命令行解析
	kingpin.Version("1.0.0")
	kingpin.Parse()

	// 全局配置
	config, err := core.NewConfig(*env)
	if err != nil {
		panic(err)
	}

	// 上下文初始化
	ctx, err := core.NewContext(config)
	if err != nil {
		panic(err)
	}

	// 服务初始化
	server := api.NewServer(config.Mode)

	// 路由初始化
	router.NewRouter(server, ctx)

	// 服务运行
	server.GraceRun(config.Addr)
}
