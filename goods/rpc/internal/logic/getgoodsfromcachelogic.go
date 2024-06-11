package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"myFlashSale/common/mysql/goods/model"
	"myFlashSale/goods/rpc/internal/svc"
	"myFlashSale/goods/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsFromCacheLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsFromCacheLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsFromCacheLogic {
	return &GetGoodsFromCacheLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsFromCacheLogic) GetGoodsFromCache(in *pb.GetGoodsFromCacheReq) (*pb.GetGoodsFromCacheResp, error) {

	goodsIds := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, id := range goodsIds {
		key := fmt.Sprintf("goods:%d", id)
		var good model.Goods
		var cachedData []byte
		// 获取缓存数据
		err := l.svcCtx.MyCache.Get(key, &cachedData)
		if err != nil {
			if l.svcCtx.MyCache.IsNotFound(err) {
				continue // 缓存未命中，继续下一个商品
			}
			return nil, err // 其他错误，直接返回错误
		}

		// 将缓存数据反序列化
		err = json.Unmarshal(cachedData, &good)
		if err != nil {
			return nil, errors.New("反序列化缓存数据失败")
		}

		// 如果找到缓存数据，立即返回
		return &pb.GetGoodsFromCacheResp{
			Id:    good.Id,
			Name:  good.Name,
			Price: good.Price,
		}, nil
	}
	return nil, errors.New("缓存中无商品")
}
