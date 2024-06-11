package logic

import (
	"context"
	"myFlashSale/common/types"
	"myFlashSale/goods/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"myFlashSale/goods/api/internal/svc"
)

type CheckStockLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckStockLogic {
	return &CheckStockLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckStockLogic) CheckStock(req *types.StockReq) (resp *types.StockResp, err error) {

	StockResp, err := l.svcCtx.GoodsServer.CheckStock(l.ctx, &pb.StockReq{})

	return &types.StockResp{ExistStock: StockResp.Stock}, nil
}
