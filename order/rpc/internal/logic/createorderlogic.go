package logic

import (
	"context"
	"errors"
	"fmt"
	"myFlashSale/common/mysql/order/model"
	"time"

	"myFlashSale/order/rpc/internal/svc"
	"myFlashSale/order/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	userid := in.Userid
	goodsid := in.Goodsid
	timeofnow := time.Now().UnixNano()
	content := fmt.Sprintf("%v_%v_%v", userid, goodsid, timeofnow)

	NewOrder := new(model.Theorder)
	NewOrder.Userid = userid
	NewOrder.Goodsid = goodsid
	NewOrder.Content = content

	_, err := l.svcCtx.Ordermodel.Insert(l.ctx, NewOrder)
	if err != nil {
		return nil, errors.New("创建订单失败")
	}
	theorder, _ := l.svcCtx.Ordermodel.FindOneByGoodsid(l.ctx, goodsid)
	theordderid := theorder.Id
	theordercontent := theorder.Content
	return &pb.CreateOrderResponse{
		Orderid: theordderid,
		Message: fmt.Sprintf("成功获取订单：%v", theordercontent),
	}, nil
}
