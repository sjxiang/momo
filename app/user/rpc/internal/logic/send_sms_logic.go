package logic

import (
	"context"

	"momo/app/user/rpc/internal/svc"
	"momo/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsLogic) SendSms(in *pb.SendSmsRequest) (*pb.SendSmsResponse, error) {
	// todo: add your logic here and delete this line
	
	logx.Info("待实现 ... ")
	
	return &pb.SendSmsResponse{}, nil
}
