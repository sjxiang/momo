package logic

import (
	"context"
	"errors"
	"strings"

	"momo/app/bff/api/internal/svc"
	"momo/app/bff/api/internal/types"
	"momo/app/user/rpc/user"
	"momo/pkg/cryptx"
	
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
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

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	
	// 请求参数处理（过滤空格、判断非空）
	req.Name = strings.TrimSpace(req.Name)
	if len(req.Name) == 0 {
		return nil, errors.New("register name cannot be empty")
	}
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, errors.New("register mobile cannot be empty")
	}
	req.Password = strings.TrimSpace(req.Password)
	if len(req.Password) == 0 {
		return nil, errors.New("register password cannot be empty")
	}
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, errors.New("verification code cannot be empty")
	}
	
	// 请求参数校验（验证码）
	err = checkVerificationCode(l.svcCtx.BizRedis, req.Mobile, req.VerificationCode)
	if err != nil {
		logx.Errorf("checkVerificationCode error: %v", err)
		return nil, err
	}

	mobile := cryptx.Encode(req.Mobile)
	// 调用 UserRPC 判断手机号是否已经注册
	u, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &user.FindByMobileRequest{
		Mobile: mobile,
	})
	if err != nil {
		logx.Errorf("FindByMobile error: %v", err)
		return nil, err
	}
	if u != nil && u.UserId > 0 {
		return nil, errors.New("mobile has registered")
	}

	password, err := cryptx.Hash(req.Password)
	if err != nil {
		logx.Errorf("Hash error: %v", err)
	}

	// 调用 UserRPC 注册用户
	regResult, err := l.svcCtx.UserRPC.Register(l.ctx, &user.RegisterRequest{
		Username: req.Name,
		Mobile:   mobile,
		Password: password,
	})
	if err != nil {
		logx.Errorf("Register error: %v", err)
		return nil, err
	}
	
	// 卸磨杀驴
	_ = delActivationCache(req.Mobile, req.VerificationCode, l.svcCtx.BizRedis)

	return &types.RegisterResponse{
		UserId: regResult.UserId, 
	}, nil
}



func checkVerificationCode(rds *redis.Redis, mobile, code string) error {
	cacheCode, err := getActivationCache(mobile, rds)
	if err != nil {
		return err
	}
	if cacheCode == "" {
		return errors.New("verification code expired")
	}
	if cacheCode != code {
		return errors.New("verification code failed")
	}

	return nil
}  