package logic

import (
	"context"
	pb3 "myFlashSale/goods/rpc/pb/pb"
	pb2 "myFlashSale/user/rpc/pb/pb"

	"myFlashSale/payment/rpc/internal/svc"
	"myFlashSale/payment/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayLogic {
	return &PayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PayLogic) Pay(in *pb.PayReq) (*pb.PayResp, error) {
	//获取商品价格
	getPriceResp, err := l.svcCtx.Goods.GetPrice(l.ctx, &pb3.GetPriceReq{
		Goodsid: in.Goodsid,
	})
	if err != nil {
		return nil, err
	}
	//获取用户存款
	checkDepositResp, err := l.svcCtx.User.Checkdeposit(l.ctx, &pb2.CheckDepositReq{Userid: in.Userid})
	if err != nil {
		return nil, err
	}
	if getPriceResp.Price > checkDepositResp.Deposit {
		return &pb.PayResp{
			Message: "您的存款不足",
		}, nil
	} else {
		_, err := l.svcCtx.User.Updatadeposit(l.ctx, &pb2.UpdataDepositReq{
			Userid:  in.Userid,
			Account: getPriceResp.Price,
		})
		findOne, err := l.svcCtx.Paymentmodel.FindOne(l.ctx, in.Paymentid)
		if err != nil {
			return nil, err
		}
		findOne.Status = 1
		findOne.Id = in.Paymentid
		findOne.Userid = in.Userid
		findOne.Goodsid = in.Goodsid
		err = l.svcCtx.Paymentmodel.Update(l.ctx, findOne)
		if err != nil {
			return nil, err
		}
		return &pb.PayResp{
			Message: "支付成功",
		}, nil
	}
}
