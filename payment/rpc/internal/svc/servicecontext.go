package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"myFlashSale/common/mysql/payment/model"
	"myFlashSale/goods/rpc/goodsservice"
	"myFlashSale/payment/rpc/internal/config"
	"myFlashSale/user/rpc/user"
)

type ServiceContext struct {
	Config       config.Config
	Paymentmodel model.PaymentModel
	User         user.User
	Goods        goodsservice.GoodsService
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:       c,
		Paymentmodel: model.NewPaymentModel(sqlx.NewMysql(c.Mysql.DataSource)),
		User:         user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		Goods:        goodsservice.NewGoodsService(zrpc.MustNewClient(c.GoodsRpcConf)),
	}
}
