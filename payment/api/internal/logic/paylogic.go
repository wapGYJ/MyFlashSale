package logic

import (
	"context"
	"myFlashSale/payment/rpc/pb/pb"

	"myFlashSale/common/types"
	"myFlashSale/payment/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayLogic {
	return &PayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayLogic) Pay(req *types.PayReq) (resp *types.PayResp, err error) {
	payresp, err := l.svcCtx.Payment.Pay(l.ctx, &pb.PayReq{
		Userid:    req.UserId,
		Paymentid: req.PaymentId,
		Goodsid:   req.GoodsId,
	})
	return &types.PayResp{
		Message: payresp.Message,
	}, nil
}
