package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"time"

	"myFlashSale/goods/rpc/internal/svc"
	"myFlashSale/goods/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsCacheLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsCacheLogic {
	return &GoodsCacheLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GoodsCacheLogic) GoodsCache(in *pb.CacheReq) (*pb.CacheResp, error) {
	// 从数据库中加载商品数据
	goods, err := l.svcCtx.GoodsModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	// 初始化 goodsCache
	goodsCache := make(map[int64][]byte)
	// 将商品数据写入 Redis 缓存
	for _, good := range goods {
		key := fmt.Sprintf("goods:%d", good.Id)
		data, err := json.Marshal(good)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal goods data")
		}
		err = l.svcCtx.MyCache.SetWithExpire(key, data, time.Hour)
		if err != nil {
			return nil, errors.Wrap(err, "failed to set cache")
		}
		goodsCache[good.Id] = data
		l.Infof("Cached data for goods with ID %d.", good.Id)
	}
	l.Infof("Preloaded %d goods into Redis cache.", len(goods))

	return &pb.CacheResp{
		Success: true,
	}, nil
}
