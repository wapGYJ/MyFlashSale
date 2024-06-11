package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	GoodsRpcConf zrpc.RpcClientConf
	Rabbitmq     struct {
		Url   string
		Queue string
	}
}
