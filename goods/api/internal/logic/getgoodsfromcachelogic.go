package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"myFlashSale/common/types"
	"myFlashSale/goods/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"myFlashSale/goods/api/internal/svc"
)

type GetGoodsFromCacheLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodsFromCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsFromCacheLogic {
	return &GetGoodsFromCacheLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsFromCacheLogic) GetGoodsFromCache(req *types.GetGoodsFromCacheReq) (resp *types.GetGoodsFromCacheResp, err error) {
	GetGoodsFromCacheResp, err := l.svcCtx.GoodsServer.GetGoodsFromCache(l.ctx, &pb.GetGoodsFromCacheReq{})
	var Resp types.GetGoodsFromCacheResp
	_ = copier.Copy(&Resp, GetGoodsFromCacheResp)
	return &Resp, nil
}
