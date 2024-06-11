package logic

import (
	"context"

	"myFlashSale/goods/rpc/internal/svc"
	"myFlashSale/goods/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPriceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPriceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPriceLogic {
	return &GetPriceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPriceLogic) GetPrice(in *pb.GetPriceReq) (*pb.GetPriceResp, error) {
	findOne, err := l.svcCtx.GoodsModel.FindOne(l.ctx, in.Goodsid)
	if err != nil {
		return nil, err
	}
	return &pb.GetPriceResp{Price: findOne.Price}, nil
}
