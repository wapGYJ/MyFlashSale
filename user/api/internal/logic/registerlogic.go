package logic

import (
	"context"
	"myFlashSale/user/rpc/pb/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"myFlashSale/common/types"
	"myFlashSale/user/api/internal/svc"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserRequest) (resp *types.UserResponse, err error) {
	RegisterResp, err := l.svcCtx.User.Register(l.ctx, &pb.UserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	var Resp types.UserResponse

	_ = copier.Copy(&Resp, RegisterResp)

	return &Resp, nil

}
