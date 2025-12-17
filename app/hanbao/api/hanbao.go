package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"hanbao-engine/app/hanbao/api/internal/config"
	"hanbao-engine/app/hanbao/api/internal/handler"
	"hanbao-engine/app/hanbao/api/internal/svc"
)

var configFile = flag.String("f", "etc/hanbao-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c.RestConf)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("🚀 汉字寻宝引擎启动成功!\n")
	fmt.Printf("📍 服务地址: http://%s:%d\n", c.Host, c.Port)
	fmt.Printf("🎯 准备开启15分钟的解谜之旅...\n")

	server.Start()
}
