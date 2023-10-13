package logic

import (
	"context"
	"encoding/json"

	"momo/app/bff/api/internal/svc"
	"momo/app/bff/api/internal/types"
	"momo/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// 提取 userId
	userId, err := l.ctx.Value(types.UserIdKey).(json.Number).Int64()
	if err != nil {
		return nil, err 
	}
	if userId == 0 {
		return nil, nil
	}

	// 调用 UserRPC 
	u, err := l.svcCtx.UserRPC.FindById(l.ctx, &user.FindByIdRequest{
		UserId: userId,
	})
	if err != nil {
		logx.Errorf("FindById userId: %d, error: %v", userId, err)
		return nil, err
	}

	return &types.UserInfoResponse{
		UserId:   u.UserId,
		Username: u.Username,
		Avatar:   u.Avatar,
	}, nil

}
