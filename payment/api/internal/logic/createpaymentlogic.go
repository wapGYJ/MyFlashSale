package logic

import (
	"context"
	"myFlashSale/payment/rpc/pb/pb"

	"myFlashSale/common/types"
	"myFlashSale/payment/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePaymentLogic {
	return &CreatePaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePaymentLogic) CreatePayment(req *types.CreatePaymentReq) (resp *types.CreatePaymentResp, err error) {
	createpaymentresp, err := l.svcCtx.Payment.CreatePayment(l.ctx, &pb.CreatePaymentRequest{
		Orderid: req.OrderId,
		Userid:  req.UserId,
		Goodsid: req.GoodsId,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreatePaymentResp{
		PaymentId: createpaymentresp.Paymentid,
		Status:    createpaymentresp.Status,
	}, nil
}
