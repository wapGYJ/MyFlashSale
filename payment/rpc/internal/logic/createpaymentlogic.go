package logic

import (
	"context"
	"errors"
	"myFlashSale/common/mysql/payment/model"

	"myFlashSale/payment/rpc/internal/svc"
	"myFlashSale/payment/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePaymentLogic {
	return &CreatePaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePaymentLogic) CreatePayment(in *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	var payment model.Payment
	payment.Goodsid = in.Goodsid
	payment.Userid = in.Userid
	payment.Status = 0
	_, err := l.svcCtx.Paymentmodel.Insert(l.ctx, &payment)
	if err != nil {
		return nil, errors.Unwrap(err)
	}
	oneByGoodsid, err := l.svcCtx.Paymentmodel.FindOneByGoodsid(l.ctx, in.Goodsid)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePaymentResponse{
		Paymentid: oneByGoodsid.Id,
		Status:    oneByGoodsid.Status,
	}, nil
}
