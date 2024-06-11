package svc

import (
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/zrpc"
	"myFlashSale/common/mq"
	"myFlashSale/user/api/internal/config"
	"myFlashSale/user/rpc/user"
)

type ServiceContext struct {
	Config    config.Config
	User      user.User
	Mq        *amqp.Channel
	QueueName string
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := mq.NewRabbitMQConnection(c.Rabbitmq.Url)
	ch := mq.NewRabbitMQChannel(conn)
	queue := mq.NewRabbitMQQueue(ch, c.Rabbitmq.Queue)

	return &ServiceContext{
		Config:    c,
		User:      user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		Mq:        ch,
		QueueName: queue.Name,
	}
}
