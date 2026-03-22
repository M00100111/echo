// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"Echo/api/internal/config"
	"Echo/api/internal/handler"
	"Echo/api/internal/svc"
	"flag"
	"fmt"
	"github.com/joho/godotenv"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/agent.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	err := godotenv.Load("./etc/.env")
	if err != nil {
		panic(err)
	}
	c.InitConfig()
	fmt.Printf("config: %+v", c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
