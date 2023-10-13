package logic

import (
	"context"

	"momo/app/user/rpc/internal/svc"
	"momo/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByMobileLogic {
	return &FindByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 填充业务逻辑
func (l *FindByMobileLogic) FindByMobile(in *pb.FindByMobileRequest) (*pb.FindByMobileResponse, error) {
	
	// todo - sqlx 查询结果处理，需要研究下
	
	user, err := l.svcCtx.UserModel.CustomFindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		logx.Errorf("FindOneByMobile mobile: %s, error: %v", in.Mobile, err)
		return nil, err
	}
	if user == nil {
		return &pb.FindByMobileResponse{}, nil
	}

	return &pb.FindByMobileResponse{
		UserId:   user.Id,
		Username: user.Username,
		Avatar:   user.Avatar,
	}, nil

}
