package logic

import (
	"context"
	"time"

	"momo/app/user/rpc/internal/model"
	"momo/app/user/rpc/internal/svc"
	"momo/app/user/rpc/pb"

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

func (l *RegisterLogic) Register(in *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	now := time.Now()
	
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username:   in.Username,
		Mobile:     in.Mobile,
		Avatar:     in.Avatar,
		CreateTime: now,
		UpdateTime: now,
	})
	if err != nil {
		logx.Errorf("Insert req: %v, err: %v", in, err)
		return nil, err
	}
	userId, err := res.LastInsertId()
	if err != nil {
		logx.Errorf("LastInsertId error: %v", err)
		return nil, err
	}

	return &pb.RegisterResponse{
		UserId: userId,
	}, nil
}
