package logic

import (
	"context"
	"myFlashSale/order/rpc/pb/pb"

	"myFlashSale/common/types"
	"myFlashSale/order/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderReq) (resp *types.CreateOrderResp, err error) {
	CreateOrderResp, err := l.svcCtx.Order.CreateOrder(l.ctx, &pb.CreateOrderRequest{
		Userid:  req.UserId,
		Goodsid: req.GoodsId,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateOrderResp{
		OrderId: CreateOrderResp.Orderid,
		Message: "创建订单成功",
	}, nil
}
