package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"myFlashSale/common/jwt"
	"myFlashSale/common/mysql/user/model"

	"myFlashSale/user/rpc/internal/svc"
	"myFlashSale/user/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.UserRequest) (*pb.UserResponse, error) {
	UsernameFromReq := in.Username

	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, UsernameFromReq)
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return nil, err
	}
	if user != nil {
		err = errors.New("用户已存在")
		return nil, err
	}
	//插入
	newuser := new(model.User)
	newuser.Username = in.Username
	newuser.Password = in.Password
	newuser.Deposit = 5000
	_, err = l.svcCtx.UserModel.Insert(l.ctx, newuser)
	if err != nil {
		return nil, errors.New("创建用户失败")
	}
	theuser, _ := l.svcCtx.UserModel.FindOneByUsername(l.ctx, UsernameFromReq)
	UserId := theuser.Id
	UserName := theuser.Username
	//签发token
	myAuth := l.svcCtx.Config.JwtAuth
	myAuthSecret := myAuth.JwtAccessSecret
	myAuthExpire := myAuth.JwtAccessExpire
	token, err := jwt.GenToken(jwt.JwtPayLoad{
		UserID:   UserId,
		Username: UserName,
		Role:     1,
	}, myAuthSecret, myAuthExpire)
	if err != nil {
		return nil, errors.New("签发token失败")
	}

	return &pb.UserResponse{
		Userid: UserId,
		Token:  token,
	}, nil

}
