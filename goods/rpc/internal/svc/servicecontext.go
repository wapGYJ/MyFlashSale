package svc

import (
	"errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
	"myFlashSale/common/mysql/goods/model"
	"myFlashSale/goods/rpc/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	GoodsModel model.GoodsModel
	MyCache    cache.Cache
	MyRedis    redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	MysqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	NewRedis, _ := redis.NewRedis(c.MyRedis)
	// 创建 SingleFlight 实例
	barrier := syncx.NewSingleFlight()
	// 初始化 Cache 对象
	cacheCluster := cache.New(c.MyCache, barrier, cache.NewStat("cache"), errors.New("cache not found"))

	return &ServiceContext{
		Config:     c,
		GoodsModel: model.NewGoodsModel(MysqlConn),
		MyCache:    cacheCluster,
		MyRedis:    *NewRedis,
	}
}
