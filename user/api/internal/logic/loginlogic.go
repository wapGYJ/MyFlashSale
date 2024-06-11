package logic

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/streadway/amqp"
	"myFlashSale/user/rpc/pb/pb"

	"myFlashSale/common/types"
	"myFlashSale/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserRequest) (resp *types.UserResponse, err error) {
	LoginResp, err := l.svcCtx.User.Login(l.ctx, &pb.UserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	var Resp types.UserResponse

	_ = copier.Copy(&Resp, LoginResp)

	if err != nil {
		return nil, err
	}
	// 将checkstock请求放入队列
	checkStockReq := &types.StockReq{}
	body, err := json.Marshal(checkStockReq)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.Mq.Publish(
		"",                 // exchange
		l.svcCtx.QueueName, // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return nil, err
	}
	return &Resp, nil
}
