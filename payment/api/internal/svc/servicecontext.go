package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"myFlashSale/payment/api/internal/config"
	"myFlashSale/payment/rpc/paymentservice"
)

type ServiceContext struct {
	Config  config.Config
	Payment paymentservice.PaymentService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Payment: paymentservice.NewPaymentService(zrpc.MustNewClient(c.PaymentRpcConf)),
	}
}
