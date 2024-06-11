package logic

import (
	"context"

	"myFlashSale/user/rpc/internal/svc"
	"myFlashSale/user/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckdepositLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckdepositLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckdepositLogic {
	return &CheckdepositLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckdepositLogic) Checkdeposit(in *pb.CheckDepositReq) (*pb.CheckDepositResp, error) {
	one, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Userid)
	if err != nil {
		return nil, err
	}
	deposit := one.Deposit
	return &pb.CheckDepositResp{
		Deposit: deposit,
	}, nil
}
