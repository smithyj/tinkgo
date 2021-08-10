package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"tinkgo/pkg/tinkgo/httpx"
	"tinkgo/pkg/tinkgo/logx"
	"tinkgo/service/app/api/internal/config"
	"tinkgo/service/app/api/internal/handler"
	"tinkgo/service/app/api/internal/svc"
)

var env = kingpin.Flag("env", "Set run environment, options: prod / test / dev").Default("dev").String()

func main() {
	// 命令行解析
	kingpin.Version("1.0.0")
	kingpin.Parse()

	// 全局配置
	c, err := config.NewConfig(*env)
	if err != nil {
		panic(err)
	}

	// 日志初始化
	logx.Setup(c.LogConfig)

	logx.Debug().Msg("1")
	logx.Info().Msg("2")
	logx.Warn().Msg("3")
	logx.Error().Msg("4")
	logx.Trace().Msg("5")
	logx.Fatal().Msg("6")
	logx.Panic().Msg("7")

	// 上下文初始化
	srvCtx, err := svc.NewServiceContext(c)
	if err != nil {
		panic(err)
	}

	// 服务初始化
	server := httpx.NewServer(c.Mode)

	// 路由初始化
	handler.NewRouter(server, srvCtx)

	// 服务运行
	server.GraceRun(c.Addr)
}
