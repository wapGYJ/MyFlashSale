package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"myFlashSale/common/mysql/order/model"
	"myFlashSale/order/rpc/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	Ordermodel model.TheorderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	MysqlConn := sqlx.NewMysql(c.Mysql.Datasource)

	return &ServiceContext{
		Config:     c,
		Ordermodel: model.NewTheorderModel(MysqlConn),
	}
}
