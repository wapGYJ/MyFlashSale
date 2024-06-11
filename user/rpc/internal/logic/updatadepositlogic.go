package logic

import (
	"context"
	"errors"
	"myFlashSale/user/rpc/internal/svc"
	"myFlashSale/user/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatadepositLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatadepositLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatadepositLogic {
	return &UpdatadepositLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatadepositLogic) Updatadeposit(in *pb.UpdataDepositReq) (*pb.UpdataDepositResp, error) {
	one, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Userid)
	if err != nil {
		return nil, err
	}
	if one.Deposit < in.Account {
		return nil, errors.New("存款不足")
	} else {
		one.Deposit -= in.Account
		err := l.svcCtx.UserModel.Update(l.ctx, one)
		if err != nil {
			return nil, errors.New("更新存款失败")
		}
		return &pb.UpdataDepositResp{Msg: "存款已更新"}, nil
	}
}
