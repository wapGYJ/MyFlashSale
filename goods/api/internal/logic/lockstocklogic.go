package logic

import (
	"context"
	"myFlashSale/common/types"
	"myFlashSale/goods/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"myFlashSale/goods/api/internal/svc"
)

type LockStockLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLockStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LockStockLogic {
	return &LockStockLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LockStockLogic) LockStock(req *types.LockStockRequest) (resp *types.LockStockResponse, err error) {
	LockStockResponse, err := l.svcCtx.GoodsServer.LockStock(l.ctx, &pb.LockStockRequest{
		GoodsId: req.GoodsId,
	})

	return &types.LockStockResponse{
		LockId:   LockStockResponse.GoodsId,
		Quantity: LockStockResponse.Quantity,
	}, nil
}
