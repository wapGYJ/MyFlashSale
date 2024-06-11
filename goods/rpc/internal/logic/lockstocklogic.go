package logic

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"myFlashSale/goods/rpc/internal/svc"
	"myFlashSale/goods/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LockStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLockStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LockStockLogic {
	return &LockStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LockStockLogic) LockStock(in *pb.LockStockRequest) (*pb.LockStockResponse, error) {
	// Redis 锁的 key 和超时时间
	lockKey := fmt.Sprintf("lock:goods:%s", strconv.FormatInt(in.GoodsId, 10))
	lockTimeout := int(5 * time.Second)

	ok, err := l.svcCtx.MyRedis.SetnxEx(lockKey, "1", lockTimeout)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("获取锁失败")
	}

	//释放锁
	defer func() {
		_, err = l.svcCtx.MyRedis.Del(lockKey)
		if err != nil {
			return
		}
	}()
	// 缓存删除操作

	stockKey := fmt.Sprintf("goods:%s", strconv.FormatInt(in.GoodsId, 10))
	luaScript := `
		if redis.call('EXISTS', KEYS[1]) == 1 then
			redis.call('DEL', KEYS[1])
			return 1
		else
			return 0
		end
	`
	result, err := l.svcCtx.MyRedis.Eval(luaScript, []string{stockKey}, "")
	if result == 0 || err != nil {
		return nil, errors.New("此商品缓存不存在或未能删除库存缓存")
	}
	return &pb.LockStockResponse{
		GoodsId:  in.GoodsId,
		Quantity: 1,
	}, nil
}
