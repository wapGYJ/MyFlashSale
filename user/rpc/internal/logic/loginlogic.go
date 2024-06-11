package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"myFlashSale/common/jwt"

	"myFlashSale/user/rpc/internal/svc"
	"myFlashSale/user/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.UserRequest) (*pb.UserResponse, error) {
	UserNameFromReq := in.Username
	PasswordFromReq := in.Password
	oneByUsername, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, UserNameFromReq)
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errors.New("查无此人，请先进行注册")
		}
		return nil, err
	}
	if oneByUsername.Password != PasswordFromReq {
		return nil, errors.New("密码错误")
	} else {
		userid := oneByUsername.Id
		username := oneByUsername.Username
		//签发token
		myAuth := l.svcCtx.Config.JwtAuth
		myAuthSecret := myAuth.JwtAccessSecret
		myAuthExpire := myAuth.JwtAccessExpire
		token, err := jwt.GenToken(jwt.JwtPayLoad{
			UserID:   userid,
			Username: username,
			Role:     1,
		}, myAuthSecret, myAuthExpire)
		if err != nil {
			return nil, errors.New("签发token失败")
		}
		return &pb.UserResponse{
			Userid: userid,
			Token:  token,
		}, nil
	}
}
