package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"myFlashSale/goods/rpc/internal/svc"
	"myFlashSale/goods/rpc/pb/pb"
)

type CheckStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckStockLogic {
	return &CheckStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckStockLogic) CheckStock(in *pb.StockReq) (*pb.StockResp, error) {

	var totalKeys int64

	// 初始化游标
	cursor := uint64(0)
	// 初始化匹配模式（这里表示获取所有的键）
	pattern := "*"

	for {
		// 使用 SCAN 命令遍历 Redis 中的键
		keys, nextCursor, err := l.svcCtx.MyRedis.Scan(cursor, pattern, 10)
		if err != nil {
			return nil, err
		}

		// 计算总键数
		totalKeys += int64(len(keys))

		// 如果 nextCursor 为 0，则表示遍历结束
		if nextCursor == 0 {
			break
		}

		cursor = nextCursor
	}
	return &pb.StockResp{Stock: totalKeys}, nil
}
