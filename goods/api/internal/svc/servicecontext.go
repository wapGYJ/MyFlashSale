package svc

import (
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/zrpc"
	"myFlashSale/common/mq"
	"myFlashSale/goods/api/internal/config"
	"myFlashSale/goods/rpc/goodsservice"
)

type ServiceContext struct {
	Config      config.Config
	GoodsServer goodsservice.GoodsService
	Mq          *amqp.Channel
	QueueName   string
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := mq.NewRabbitMQConnection(c.Rabbitmq.Url)
	ch := mq.NewRabbitMQChannel(conn)
	queue := mq.NewRabbitMQQueue(ch, c.Rabbitmq.Queue)
	return &ServiceContext{
		Config:      c,
		GoodsServer: goodsservice.NewGoodsService(zrpc.MustNewClient(c.GoodsRpcConf)),
		Mq:          ch,
		QueueName:   queue.Name,
	}
}
