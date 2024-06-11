package main

import (
	"flag"
	"fmt"
	"myFlashSale/goods/api/internal/config"
	"myFlashSale/goods/api/internal/handler"
	"myFlashSale/goods/api/internal/mq"
	"myFlashSale/goods/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "D:\\golang--work\\myFlashSale\\goods\\api"+
	"\\etc\\goods-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	mq.ConsumeCheckStock(ctx, ctx.Mq, ctx.QueueName)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
