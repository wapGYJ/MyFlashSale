package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"myFlashSale/order/api/internal/config"
	"myFlashSale/order/rpc/orderservice"
)

type ServiceContext struct {
	Config config.Config
	Order  orderservice.OrderService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Order:  orderservice.NewOrderService(zrpc.MustNewClient(c.OrderRpcConf)),
	}
}
